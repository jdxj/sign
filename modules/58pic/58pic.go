package _8pic

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	config "sign/utils/conf"
	"sign/utils/log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func New58PicFromApi(conf *config.Pic58Conf) (*Toucher58pic, error) {
	if conf == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &Toucher58pic{
		name:               conf.Name,
		cookies:            conf.Cookies,
		loginURL:           "https://www.58pic.com/index.php?m=IntegralMall",
		verifyKey:          ".cs-ul3-li1",
		verifyReverseValue: "我的积分:--",
		signDataURL:        "https://www.58pic.com/index.php?m=jifenNew&a=getTreeActivity",
		signURL:            "https://www.58pic.com/index.php?m=signin&a=addUserSign&time=",
		client:             &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

type Toucher58pic struct {
	name string
	// 用 "key=value; key=value" 表示的 cookie 字符串,
	// 其主要用于第一次启动所使用的 cookie, 登录成功后使用 http.Client 管理.
	cookies            string
	loginURL           string // 用于验证是否登录成功所要抓取的网页
	verifyKey          string // 指定要抓取得属性, 比如 class, li 等 html 标签或属性
	verifyReverseValue string // 当要抓取的属性等于 VerifyValue 时, 判断为登录失败
	signDataURL        string // 执行签到签获取签到数据的链接
	signURL            string // 执行签到所要访问的链接

	client *http.Client

	// 模拟浏览用
	loginStat bool
	browsing  bool
}

func (tou *Toucher58pic) Name() string {
	return tou.name
}

func (tou *Toucher58pic) Boot() bool {
	cookies, err := utils.StrToCookies(tou.cookies, utils.Pic58CookieDomain)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}

	cookieURL, err := url.Parse(utils.Pic58CookieURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

// Login 58pic 的登录使用 cookie 方式
func (tou *Toucher58pic) Login() bool {
	resp, err := tou.client.Get(tou.loginURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}

	var mark bool
	doc.Find(tou.verifyKey).Each(func(i int, selection *goquery.Selection) {
		if !strings.HasSuffix(selection.Text(), "--") {
			mark = true
		} else {
			log.MyLogger.Info("%s redeem info not found", log.Log_58pic)
		}
	})

	tou.loginStat = mark
	tou.mockBrowsing()
	return mark
}

// 签到前所需数据
type conf struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Status int `json:"status"`
		Msg    struct {
			CycleTime string `json:"cycle_time"`
		} `json:"msg"`
	} `json:"data"`
}

// 签到后要验证的数据
type sign struct {
	Status      string `json:"status"`
	Type        string `json:"type"`
	Times       string `json:"times"`
	ClickNum    int    `json:"clickNum"`
	Week        string `json:"week"`
	RewardThing string `json:"rewardThing"`
}

func (tou *Toucher58pic) Sign() bool {
	val := url.Values{
		"taskIdNum": []string{"40"},
	}
	resp, err := tou.client.PostForm(tou.signDataURL, val)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}

	conf := &conf{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		log.MyLogger.Error("%s %s data: %s", log.Log_58pic, err, data)
		return false
	}

	cycTime := conf.Data.Msg.CycleTime
	s, e := beginAndEnd()
	signURL := tou.signURL + fmt.Sprintf("%d", utils.NowUnixMilli())

	val = url.Values{
		"cycle":      []string{cycTime},
		"sign":       []string{},
		"start_time": []string{s},
		"end_time":   []string{e},
	}
	resp, err = tou.client.PostForm(signURL, val)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_58pic, err)
		return false
	}

	sign := &sign{}
	err = json.Unmarshal(data, sign)
	if err != nil {
		log.MyLogger.Error("%s %s data: %s", log.Log_58pic, err, data)
		return false
	}

	if sign.Status == "1" {
		return true
	}
	return false
}

func (tou *Toucher58pic) mockBrowsing() {
	if tou.browsing {
		return
	}
	tou.browsing = true

	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		defer log.MyLogger.Info("%s mock browsing finish", log.Log_58pic)

		for tou.loginStat {
			<-ticker.C

			req, _ := http.NewRequest("GET", "https://www.58pic.com/", nil)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
			tou.client.Do(req)
		}
	}()
}

func beginAndEnd() (string, string) {
	t := time.Now()
	y, mNamed, d := t.Date()
	m := int(mNamed)
	w := int(t.Weekday())

	// 0-6: 星期日->星期三->星期六
	// 星期一是多少号
	weekday1 := d - (w - 1)
	// 星期日是多少号
	weekday0 := weekday1 + 6

	// 年 月 日
	data := fmt.Sprintf("%d-%s-%s", y, fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", weekday1))
	start := base64.StdEncoding.EncodeToString([]byte(data))

	data = fmt.Sprintf("%d-%s-%s", y, fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", weekday0))
	end := base64.StdEncoding.EncodeToString([]byte(data))
	return start, end
}
