package bilibili_com

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"sign/utils"
)

// todo: 签到失败时不要 panic
// 硬币查询链接
// https://account.bilibili.com/account/coin
func Start() {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	err := c.SetCookies("https://www.bilibili.com", utils.Cookies("bilibili.com", utils.Cookie_Bilibili))
	if err != nil {
		utils.MyLogger.Error("%s %s", utils.Log_Bilibili, err)
		return
	}

	//c.OnResponse(func(response *colly.Response) {
	//	fmt.Println("Header:", response.Headers)
	//	fmt.Println("StatusCode:", response.StatusCode)
	//	fmt.Printf("Body: %s", response.Body)
	//})

	utils.MyLogger.Info("%s %s", utils.Log_Bilibili, "访问主页")
	c.Visit("https://www.bilibili.com")
}
