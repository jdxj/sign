package pic

import (
	"encoding/base64"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"sign/utils"
	"sync"
	"time"
)

// 签到链接
// /index.php?m=signin&a=addUserSign&time=1570551897063
// form 的格式: (其中的 start_time/end_time 是经过 base64 加密的从周一到周日的日期, 例如:2019-10-07)
// cycle: 100
// sign:
// start_time: MjAxOS0xMC0wNw==
// end_time: MjAxOS0xMC0xMw==

func Start() {
	// once 用于只发送一次签到行动
	// 因为在查找 html 元素会有多个相同目标元素
	var once sync.Once
	var cycleTime string

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	err := c.SetCookies("https://www.58pic.com", cookies())
	if err != nil {
		utils.MyLogger.Error("%s %s", utils.Log_58pic, err)
		return
	}

	// 获取 cycle_time
	c.OnResponse(func(resp *colly.Response) {
		buf := make([]byte, len(resp.Body))
		copy(buf, resp.Body)
		ct := CycleTime(buf)
		if ct != "" {
			cycleTime = ct
		}
	})

	// todo: 成功登录的 html class 会因为千图页面的更改而随之更改
	c.OnHTML(".user-info", func(e *colly.HTMLElement) {
		once.Do(func() {
			// 打印用户 ID
			utils.MyLogger.Info("%s %s", utils.Log_58pic, e.Text)
			utils.MyLogger.Info("%s %s", utils.Log_58pic, "获取 cycle_time")
			c.Post(cycleTimeUrl(), cycleTimeData())

			utils.MyLogger.Info("%s %s", utils.Log_58pic, "执行签到")
			// 访问签到链接
			c.Post(postUrl(), postData(cycleTime))
		})

		// 访问积分明细页
		c.Visit("https://www.58pic.com/index.php?m=IntegralMall&a=qtbRecord")
	})

	// 获取积分收支明细
	c.OnHTML(".szmx-list", func(e *colly.HTMLElement) {
		// 打印积分
		utils.MyLogger.Info("%s %s", utils.Log_58pic, e.Text)
	})

	//c.Visit("https://www.58pic.com/")
	c.Visit("https://www.58pic.com/index.php?m=IntegralMall")
}

// todo: 根据响应头更新 cookies
func cookies() []*http.Cookie {
	var cookies []*http.Cookie

	kvs := utils.ConfAll("58pic.com")
	// 无所谓的过期时间
	expires := time.Date(2048, 1, 1, 0, 0, 0, 0, time.Now().Location())
	for _, kv := range kvs {
		cookie := &http.Cookie{
			Name:     kv.K,
			Value:    kv.V,
			Path:     "/",
			Domain:   ".58pic.com",
			Expires:  expires,
			Secure:   false,
			HttpOnly: false,
		}

		cookies = append(cookies, cookie)
	}

	if len(cookies) != 0 {
		utils.MyLogger.Debug("%s %s", utils.Log_58pic, "读取配置成功")
	}
	return cookies
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

func unixTimeMill() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
}

func postUrl() string {
	url := "https://www.58pic.com/index.php?m=signin&a=addUserSign&time=" + unixTimeMill()
	utils.MyLogger.Info("%s 签到 url: %s", utils.Log_58pic, url)
	return url
}

func postData(ct string) map[string]string {
	m := make(map[string]string)
	s, e := beginAndEnd()

	m["cycle"] = ct
	m["sign"] = ""
	m["start_time"] = s
	m["end_time"] = e

	utils.MyLogger.Info("%s 签到 map: %s", utils.Log_58pic, m)
	return m
}

func cycleTimeUrl() string {
	url := "https://www.58pic.com/index.php?m=jifenNew&a=getTreeActivity"
	utils.MyLogger.Info("%s ct url: %s", utils.Log_58pic, url)
	return url
}

func cycleTimeData() map[string]string {
	m := make(map[string]string)
	m["taskIdNum"] = "40"

	return m
}
