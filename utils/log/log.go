package log

import (
	"github.com/astaxie/beego/logs"
)

const (
	Log_StudyGolang = "[studygolang]"
	Log_Bilibili    = "[bilibili]"
	Log_58pic       = "[58pic]"
	Log_HacPai      = "[hacpai]"
	Log_V2ex        = "[v2ex]"
	Log_Main        = "[main]"
	Log_Log         = "[log]"
	Log_Cmd         = "[cmd]"
	Log_Task        = "[task]"
)

var (
	MyLogger *logs.BeeLogger
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"sign.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	MyLogger = logs.GetBeeLogger()
}
