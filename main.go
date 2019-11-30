package main

import (
	"sign/modules/service"
	"sign/utils/conf"
	"sign/utils/email"
)

func main() {
	sec := conf.Conf.Section("email")
	username := sec.Key("username").String()
	msg := &email.Msg{
		To:      username,
		Subject: email.SignSetup,
		Content: "Sign 程序已启动, 注意日志和邮件",
	}
	email.SendEmail(msg)

	// 利用 web 的 listenXXX() 来阻塞 main()
	service.Service()
}
