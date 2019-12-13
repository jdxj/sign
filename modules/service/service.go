package service

import (
	"net/http"
	"sign/utils/conf"
	"sign/utils/log"

	"github.com/gin-gonic/gin"
)

func Service() {
	basicAuth := conf.Conf.Section("basicauth")
	if basicAuth == nil {
		panic("not found basicauth section in config file")
	}
	baUsername := basicAuth.Key("username").String()
	baPassword := basicAuth.Key("password").String()

	engine := gin.New()
	engine.Use(Logger)

	engine.GET("/ping", Pong)

	apiRouter := engine.Group("/api")
	apiRouter.Use(gin.BasicAuth(gin.Accounts{
		baUsername: baPassword,
	}))
	{ // service
		apiRouter.GET("/ping", ApiPong)
		apiRouter.POST("/studygolang", SignStudyGolang)
		apiRouter.POST("/bilibili", SignBili)
		apiRouter.POST("/58pic", Sign58Pic)
		apiRouter.POST("/hacpai", SignHacPai)
		apiRouter.POST("/v2ex", SignV2ex)
		apiRouter.POST("/iqiyi", SignIQiYi)
	}
	{ // manage
		apiRouter.GET("/listTask", ListTask)
	}

	engine.Run(":49159")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func Logger(c *gin.Context) {
	log.MyLogger.Info("%s %s", log.Log_API, c.Request.URL)
}
