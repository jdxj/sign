package studygolang_com

import (
	"fmt"
	"github.com/gocolly/colly"
)

// 登录: https://studygolang.com/account/login POST
// 登出: https://studygolang.com/account/logout GET

// 任务页面: https://studygolang.com/mission/daily GET
// 点击签到: https://studygolang.com/mission/daily/redeem GET

func Start() {
	c := colly.NewCollector(
		colly.AllowedDomains("studygolang.com"),
	)
	// todo: 登录参数隐藏
	if err := c.Post("https://studygolang.com/account/login", map[string]string{
		"redirect_uri": "https://studygolang.com/",
		"username":     "985759262@qq.com",
		"passwd":       "",
	}); err != nil {
		panic(err)
	}
	fmt.Println("登录成功!")

	c.OnHTML(".user-name", func(e *colly.HTMLElement) {
		fmt.Println("user name:", e.Text)
	})

	c.Visit("https://studygolang.com/")
}
