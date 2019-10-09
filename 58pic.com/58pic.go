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

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	err := c.SetCookies("https://www.58pic.com", cookies())
	if err != nil {
		utils.LogPrintln(utils.Log_58pic, err)
		return
	}

	// .user-info: 登录成功后才有
	c.OnHTML(".user-info", func(e *colly.HTMLElement) {
		once.Do(func() {
			// 打印用户 ID
			utils.LogPrintln(utils.Log_58pic, e.Text)
			utils.LogPrintln(utils.Log_58pic, "执行签到")
			c.Post(postUrl(), postData())
		})

		// 以下语句不能放到 once 中, 会阻塞, 原因不明, 但是源头在底层的 goquery
		// colly 会排除重复 url, 先写在这里
		// 访问积分页
		c.Visit("https://www.58pic.com/index.php?m=IntegralMall")
	})

	// .cs-ul3-li1: 查询积分页
	c.OnHTML(".cs-ul3-li1", func(e *colly.HTMLElement) {
		// 打印积分
		utils.LogPrintln(utils.Log_58pic, e.Text)
	})

	c.Visit("https://www.58pic.com/")
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
		utils.LogPrintln(utils.Log_58pic, "读取配置成功")
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
	utils.LogPrintln(utils.Log_58pic, "签到 url:", url)
	return url
}

func postData() map[string]string {
	m := make(map[string]string)
	s, e := beginAndEnd()

	m["cycle"] = "100"
	m["sign"] = ""
	m["start_time"] = s
	m["end_time"] = e

	utils.LogPrintln(utils.Log_58pic, "签到 map:", m)
	return m
}
