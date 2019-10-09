package bilibili_com

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"sign/utils"
	"time"
)

// todo: 签到失败时不要 panic
// 硬币查询链接
// https://account.bilibili.com/account/coin
func Start() {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	err := c.SetCookies("https://www.bilibili.com", cookies())
	if err != nil {
		utils.LogPrintln(utils.Log_Bilibili, err)
		return
	}

	//c.OnResponse(func(response *colly.Response) {
	//	fmt.Println("Header:", response.Headers)
	//	fmt.Println("StatusCode:", response.StatusCode)
	//	fmt.Printf("Body: %s", response.Body)
	//})

	utils.LogPrintln(utils.Log_Bilibili, "访问主页")
	err = c.Visit("https://www.bilibili.com")
	if err != nil {
		utils.LogPrintln(utils.Log_Bilibili, err)
	}
}

// todo: 复用 cookies
func cookies() []*http.Cookie {
	var cookies []*http.Cookie

	kvs := utils.ConfAll("bilibili.com")
	// 无所谓的过期时间
	expires := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Now().Location())
	for _, kv := range kvs {
		cookie := &http.Cookie{
			Name:     kv.K,
			Value:    kv.V,
			Path:     "/",
			Domain:   ".bilibili.com",
			Expires:  expires,
			Secure:   false,
			HttpOnly: false,
		}

		cookies = append(cookies, cookie)
	}

	if len(cookies) != 0 {
		utils.LogPrintln(utils.Log_Bilibili, "读取配置成功")
	}
	return cookies
}
