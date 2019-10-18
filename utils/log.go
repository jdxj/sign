package utils

import (
	"github.com/astaxie/beego/logs"
)

const (
	Log_StudyGolang = "[studygolang.com]"
	Log_Bilibili    = "[bilibili.com]"
	Log_58pic       = "[58pic.com]"
)

var (
	MyLogger *logs.BeeLogger
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"sign.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	MyLogger = logs.GetBeeLogger()
}
