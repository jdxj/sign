package task

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/storage"
	"github.com/jdxj/sign/pkg/toucher"
)

var (
	tplSignInSuccess = `%s 在域 %s 签到成功`
	tplSignInFailed  = `%s 在域 %s 签到失败`
)

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Run() {
	c := cron.New()
	//_, err := c.AddFunc("0 8 * * *", testCmd)
	_, err := c.AddFunc("0 8 * * *", cmd)
	handleErr(err)
	c.Run()
}

func cmd() {
	ds := storage.Default
	uds := ds.GetAllUserData()

	for _, ud := range uds {
		retry(ud)
	}
}

func retry(val toucher.Validator) {
	var (
		count = 3

		err  error
		text string
	)

	for i := 0; i < count; i++ {
		err = val.SignIn()
		if err != nil {
			continue
		}
		err = val.Verify()
		if err != nil {
			continue
		}
		break
	}

	if err == nil {
		text = fmt.Sprintf(tplSignInSuccess,
			val.ID(), val.Domain())
	} else {
		text = fmt.Sprintf(tplSignInFailed+", err: %s",
			val.ID(), val.Domain(), err)
	}

	err = bot.Send(text)
	if err != nil {
		logger.Errorf("send %s err: %s", text, err)
	}
}

func testCmd() {
	err := bot.Send("test send message")
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	log.Printf("send ok!\n")
}
