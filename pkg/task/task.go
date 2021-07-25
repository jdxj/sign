package task

import (
	"log"

	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/pkg/storage"
)

func Run() {
	c := cron.New()
	//_, err := c.AddFunc("0 8 * * *", testCmd)
	_, err := c.AddFunc("* * * * *", testCmd)
	if err != nil {
		// todo: 通知
	}
	c.Run()
}

func cmd() {
	ds := storage.Default
	uds := ds.GetAllUserData()

	for num, ud := range uds {
		err := ud.Execute()
		if err == nil {
			// todo: 通知
			continue
		}

		ds.DelUserData(num)
		// todo: 通知
	}
}

func testCmd() {
	err := bot.Send("test send message")
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	log.Printf("send ok!\n")
}
