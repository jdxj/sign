package email

import (
	"fmt"
	"net/smtp"
	"sign/utils/conf"
	"sign/utils/log"

	"github.com/jordan-wright/email"
)

const (
	TimeFormat    = "2006-01-02 15:04:05"
	FlashActivity = "刷活跃度"
	SignSetup     = "Sign setup"
	SignStart     = "签到任务启动"
	SignSuccess   = "签到成功"
	SignFailed    = "签到失败"
)

type Msg struct {
	To string

	Subject string
	Content string
}

func SendEmail(msg *Msg) {
	sec := conf.Conf.Section("email")
	username := sec.Key("username").String()
	password := sec.Key("password").String()

	e := email.NewEmail()
	e.From = fmt.Sprintf("sign <%s>", username)
	e.To = []string{msg.To}
	e.Subject = msg.Subject
	e.Text = []byte(msg.Content)

	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", username, password, "smtp.qq.com"))
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Email, err)
	}
}
