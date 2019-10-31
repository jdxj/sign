package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(sub, msg string) error {
	vs := Conf("email", "username", "password")
	if vs[0] == "" || vs[1] == "" {
		return fmt.Errorf("email config error: %s", "username or password can not empty!")
	}

	e := email.NewEmail()
	e.From = fmt.Sprintf("sign <%s>", vs[0])
	e.To = []string{vs[0]}
	e.Subject = sub
	e.Text = []byte(msg)

	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", vs[0], vs[1], "smtp.qq.com"))
}
