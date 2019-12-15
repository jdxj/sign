package main

import (
	"os"
	"os/signal"
	"sign/modules/service"
	"sign/modules/task"
	"sign/utils/conf"
	"sign/utils/email"
	"sign/utils/log"
	"syscall"
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

	go service.Service()

	// 在结束前发送邮件通知
	done := make(chan os.Signal, 2)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
		task.DefaultExe.NotifyStop()
	}

	log.MyLogger.Info("%s sign stop", log.Log_Main)
}
