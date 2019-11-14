package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"sign/modules/service/static"
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
	{
		apiRouter.POST("/studygolang", SignStudyGolang)
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
