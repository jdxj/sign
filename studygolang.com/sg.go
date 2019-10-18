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
		utils.MyLogger.Error("%s %s", utils.Log_StudyGolang, err)
		return
	}

	utils.MyLogger.Info("%s %s", utils.Log_StudyGolang, "登录成功")

	c.OnHTML("p[class=c9]", func(e *colly.HTMLElement) {
		utils.MyLogger.Info("%s 领取状态: %s", utils.Log_StudyGolang, e.Text)
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

	utils.MyLogger.Debug("%s %s", utils.Log_StudyGolang, "读取配置成功")
	return res
}
