package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.String(http.StatusOK, "username: %s, password: %s, activeURL: %s",
		sgConf.Username,
		sgConf.Password,
		sgConf.ActiveURL,
	)
}
