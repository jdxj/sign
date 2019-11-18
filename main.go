package main

import (
	"sign/modules/service"
	"sign/utils/email"
	"sign/utils/log"
)

func main() {
	err := email.SendEmail("sign program start", "please notice %s file", "sign.log")
	if err != nil {
		log.MyLogger.Info("%s %s", log.Log_Main, err)
	}

	// 利用 web 的 listenXXX() 来阻塞 main()
	service.Service()
}
