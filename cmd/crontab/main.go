package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/crontab/server"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func main() {
	srv := &server.Service{}
	s := grpc.NewServer()
	crontab.RegisterTestServiceServer(s, srv)
	l, err := net.Listen("tcp", ":49152")
	if err != nil {
		panic(err)
	}
	err = s.Serve(l)
	if err != nil {
		log.Fatalln(err)
	}
}
