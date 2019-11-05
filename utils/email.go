package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"sign/utils/conf"
)

func SendEmail(sub, msg string) error {
	sec := conf.Conf.Section("email")
	username := sec.Key("username").String()
	password := sec.Key("password").String()

	e := email.NewEmail()
	e.From = fmt.Sprintf("sign <%s>", username)
	e.To = []string{username}
	e.Subject = sub
	e.Text = []byte(msg)

	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", username, password, "smtp.qq.com"))
}
