package studygolang

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils/conf"
	"sign/utils/email"
	"sign/utils/log"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func NewSGFromAPI(conf *conf.SGConf) (*ToucherStudyGolang, error) {
	if conf == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &ToucherStudyGolang{
		conf:      conf,
		loginURL:  "https://studygolang.com/account/login",
		signURL:   "https://studygolang.com/mission/daily/redeem",
		verifyKey: ".balance_area",
		signKey:   ".c9",
		signValue: "每日登录奖励已领取",
		client:    &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

type ToucherStudyGolang struct {
	conf *conf.SGConf

	loginURL  string
	signURL   string
	verifyKey string
	signKey   string
	signValue string
	client    *http.Client

	// 状态相关
	bootStat  bool // 引导是否成功
	loginStat bool // 登录是否成功
	signStat  bool // 签到是否成功

}

func (tou *ToucherStudyGolang) Name() string {
	return tou.conf.Name
}

func (tou *ToucherStudyGolang) Email() string {
	return tou.conf.To
}

func (tou *ToucherStudyGolang) Boot() bool {
	val := url.Values{
		"redirect_uri": []string{"https://studygolang.com/"},
		"username":     []string{tou.conf.Username},
		"passwd":       []string{tou.conf.Password},
		"remember_me":  []string{"1"},
	}
	resp, err := tou.client.PostForm(tou.loginURL, val)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_StudyGolang, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_StudyGolang, err)
		return false
	}

	var mark bool
	doc.Find(tou.verifyKey).Each(func(i int, selection *goquery.Selection) {
		mark = true
	})

	tou.bootStat = mark
	return mark
}

func (tou *ToucherStudyGolang) Login() bool {
	tou.loginStat = tou.bootStat
	return tou.bootStat
}

func (tou *ToucherStudyGolang) Sign() bool {
	resp, err := tou.client.Get(tou.signURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_StudyGolang, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_StudyGolang, err)
		return false
	}

	var mark bool
	doc.Find(tou.signKey).Each(func(i int, selection *goquery.Selection) {
		// 只要有一个相等就判为签到成功
		if selection.Text() == tou.signValue {
			mark = true
		}
	})
	tou.signStat = mark

	tou.run()
	return mark
}

// active 用于刷活跃度
func (tou *ToucherStudyGolang) active() {
	// 避免无限刷
	const flashLimit = 500
	count := 0

	expected := tou.conf.Expected
	if expected == 0 {
		// 随机
		expected = rand.Intn(conf.ExpectedLimit) + 1
	}
	realRanking := 10000

	for tou.signStat && count < flashLimit {
		resp, err := tou.client.Get(tou.conf.ActiveURL)
		if err != nil {
			log.MyLogger.Debug("%s execute active fail: %s", log.Log_StudyGolang, err)
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.MyLogger.Debug("%s execute active fail: %s", log.Log_StudyGolang, err)
			resp.Body.Close()
			return
		}

		doc.Find(".userinfo").Find("li").Each(func(i int, selection *goquery.Selection) {
			if i != 4 {
				return
			}

			target := selection.Find("a").Text()
			actRank, err := strconv.Atoi(target)
			if err != nil {
				log.MyLogger.Warn("%s can not parse activity ranking: %s", log.Log_StudyGolang, target)
				return
			}

			realRanking = actRank
		})
		resp.Body.Close()

		if realRanking <= expected {
			break
		}

		count++
		// 2s 刷一次
		time.Sleep(2 * time.Second)
	}

	log.MyLogger.Info("%s flash activity ranking finish, final ranking: %d", log.Log_StudyGolang, realRanking)
	log.MyLogger.Debug("%s exit activity ranking - signStat is: %s", log.Log_StudyGolang, tou.signStat)
}

// run 用于执行类似天执行一次的任务, 非阻塞的.
func (tou *ToucherStudyGolang) run() {
	// 当天21点刷活跃度
	now := time.Now()

	today0AM := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	today21PM := today0AM.Add(21 * time.Hour)

	dur := today21PM.Sub(now)
	go func() {
		if tou.conf.To != "" {
			defer func() {
				msg := &email.Msg{
					To:      tou.conf.To,
					Subject: email.FlashActivity,
					Content: fmt.Sprintf("刷 %s 的活跃度已完成, 请到官网查看结果", log.Log_StudyGolang),
				}
				email.SendEmail(msg)
			}()
		}
		// 立即执行
		if dur <= 0 {
			tou.active()
			return
		}

		timer := time.NewTimer(dur)
		<-timer.C

		tou.active()
		timer.Stop()
	}()
}
