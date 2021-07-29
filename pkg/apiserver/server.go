package apiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(host, port string) error {
	r := gin.Default()
	register(r)
	addr := fmt.Sprintf("%s:%s", host, port)
	return r.Run(addr)
}
