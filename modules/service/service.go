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
	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "/index2.html", nil)
	})

	r.Run(":49152")
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
