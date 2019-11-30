package email

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	msg := &Msg{
		To:      "985759262@qq.com",
		Subject: "test",
		Content: "ok",
	}
	SendEmail(msg)
}
