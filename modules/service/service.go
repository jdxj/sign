package service

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"sign/modules/service/static"
	"sign/utils/conf"
	"sign/utils/log"
	"strings"

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

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	engine.SetHTMLTemplate(t)
	engine.GET("/ping", Pong)

	apiRouter := engine.Group("/api")
	apiRouter.Use(gin.BasicAuth(gin.Accounts{
		baUsername: baPassword,
	}))
	{
		apiRouter.GET("/ping", ApiPong)
		apiRouter.POST("/studygolang", SignStudyGolang)
		apiRouter.POST("/bilibili", SignBili)
		apiRouter.POST("/58pic", Sign58Pic)
		apiRouter.POST("/hacpai", SignHacPai)
		apiRouter.POST("/v2ex", SignV2ex)
	}

	// todo: 可视化网页
	webRouter := engine.Group("/index")
	{
		_ = webRouter
	}

	engine.Run(":49159")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range static.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func Logger(c *gin.Context) {
	log.MyLogger.Info("%s %s", log.Log_API, c.Request.URL)
}
