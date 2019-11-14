package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"sign/modules/service/static"
	"sign/utils/log"
	"strings"
)

func Service() {
	engine := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	engine.SetHTMLTemplate(t)
	engine.GET("/ping", Pong)

	apiRouter := engine.Group("/api")
	apiRouter.Use(logger("%s somebody access %s", log.Log_API, log.Log_StudyGolang))
	{
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

	engine.Run(":49152")
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

func logger(format string, value ...interface{}) gin.HandlerFunc {
	log.MyLogger.Debug(format, value...)

	return func(context *gin.Context) {
		// 空实现, 只是为了上面的日志记录
	}
}
