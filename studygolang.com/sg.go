package studygolang_com

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"sign/utils"
)

// 登录: https://studygolang.com/account/login POST
// 登出: https://studygolang.com/account/logout GET

// 任务页面: https://studygolang.com/mission/daily GET
// 点击签到: https://studygolang.com/mission/daily/redeem GET

func Start() {
	res := conf()

	c := colly.NewCollector(
		colly.AllowedDomains("studygolang.com"),
	)
	extensions.RandomUserAgent(c)

	if err := c.Post("https://studygolang.com/account/login", map[string]string{
		"redirect_uri": "https://studygolang.com/",
		"username":     res[0],
		"passwd":       res[1],
	}); err != nil {
		panic(err)
	}

	utils.LogPrintln(utils.Log_StudyGolang, "login success!")

	c.OnHTML("p[class=c9]", func(e *colly.HTMLElement) {
		utils.LogPrintln(utils.Log_StudyGolang, "领取状态:", e.Text)
	})

	c.Visit("https://studygolang.com/mission/daily/redeem")
}

// 配置文件格式
// [studygolang.com]
//
// username = xxx
// passwd = xxx
func conf() []string {
	res := utils.Conf(utils.Conf_StudyGolang, "username", "passwd")
	if res == nil {
		panic(fmt.Errorf("%s", "user info not found"))
	}

	utils.LogPrintln(utils.Log_StudyGolang, "read conf success!")
	return res
}
