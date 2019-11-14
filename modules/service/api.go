package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sign/modules/studygolang"
	"sign/modules/task"
	"sign/utils/conf"
)

func SignStudyGolang(c *gin.Context) {
	var sgConf conf.SGConf
	if err := c.Bind(&sgConf); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "can not read data",
		})
		return
	}

	if !sgConf.CheckValidity() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "has empty data",
		})
		return
	}

	// todo: 发送数据给 Executor
	tou, err := studygolang.NewSGFromAPI(&sgConf)
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
