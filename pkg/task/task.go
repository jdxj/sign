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
	_, err := c.AddFunc("0 8 * * *", cmdSign)
	handleErr(err)
	c.Run()
}

func cmdSign() {
	ds := storage.Default
	uds := ds.GetAllUserData()

	for num, ud := range uds {
		keep := retry(ud)
		if !keep {
			ds.DelUserData(num)
		}
	}
}

func retry(val toucher.Validator) bool {
	var (
		count = 3

		keep = true
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
		keep = false
	}

	err = bot.Send(text)
	if err != nil {
		logger.Errorf("send %s err: %s", text, err)
	}
	return keep
}

func testCmd() {
	err := bot.Send("test send message")
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	log.Printf("send ok!\n")
}
