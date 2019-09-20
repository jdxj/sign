package utils

import (
	"log"
	"os"
)

const (
	Log_StudyGolang = "[studygolang.com] "
)

func LogPrintln(prefix string, v ...interface{}) {
	file, err := os.OpenFile("sign.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer file.Sync()

	logger := log.New(file, prefix, log.LstdFlags)
	logger.Println(v)
}
