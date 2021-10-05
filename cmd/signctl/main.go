package main

import (
	"log"

	"github.com/jdxj/sign/internal/signctl"
)

func main() {
	err := signctl.Execute()
	if err != nil {
		log.Println(err)
	}
}
