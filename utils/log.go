package utils

import (
	"log"
	"os"
)

const (
	Log_StudyGolang = "[studygolang.com] "
	Log_Bilibili    = "[bilibili.com] "
	Log_58pic       = "[58pic.com] "
)

// todo: 更改写入方式, 比如多长时间后关闭, 目前打开文件太频繁
func LogPrintln(prefix string, v ...interface{}) {
	file, err := os.OpenFile("sign.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer file.Sync()

	logger := log.New(file, prefix, log.LstdFlags)
	logger.Println(v...)
}
