package service

import (
	"fmt"
	"net/http"
	pic "sign/modules/58pic"
	"sign/modules/bilibili"
	"sign/modules/hacpai"
	"sign/modules/iqiyi"
	"sign/modules/studygolang"
	"sign/modules/task"
	"sign/modules/v2ex"
	"sign/utils/conf"

	"github.com/gin-gonic/gin"
)

func ApiPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func SignStudyGolang(c *gin.Context) {
	var cfg conf.SGConf
	if err := c.Bind(&cfg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !cfg.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	tou, err := studygolang.NewSGFromAPI(&cfg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	if err = task.DefaultExe.AddTaskFromApi(tou); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "add success",
	})
}

func SignBili(c *gin.Context) {
	var cfg conf.BiliConf
	if err := c.Bind(&cfg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !cfg.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	tou, err := bilibili.NewBiliFromApi(&cfg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	if err = task.DefaultExe.AddTaskFromApi(tou); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "add success",
	})
}

func Sign58Pic(c *gin.Context) {
	var cfg conf.Pic58Conf
	if err := c.Bind(&cfg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !cfg.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	tou, err := pic.New58PicFromApi(&cfg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	if err = task.DefaultExe.AddTaskFromApi(tou); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "add success",
	})
}

func SignHacPai(c *gin.Context) {
	var cfg conf.HacPaiConf
	if err := c.Bind(&cfg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !cfg.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	tou, err := hacpai.NewHacPaiFromApi(&cfg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	if err = task.DefaultExe.AddTaskFromApi(tou); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "add success",
	})
}

func SignV2ex(c *gin.Context) {
	var cfg conf.V2exConf
	if err := c.Bind(&cfg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !cfg.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	tou, err := v2ex.NewV2exFromApi(&cfg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	if err = task.DefaultExe.AddTaskFromApi(tou); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "add success",
	})
}

func SignIQiYi(c *gin.Context) {
	var cfg conf.IQiYiConf
	if err := c.Bind(&cfg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !cfg.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	tou, err := iqiyi.NewIQiYiFromApi(&cfg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	if err = task.DefaultExe.AddTaskFromApi(tou); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "add success",
	})
}

func ListTask(c *gin.Context) {
	names := task.DefaultExe.GetTaskNames()

	c.JSON(http.StatusOK, gin.H{
		"names": names,
	})
}
