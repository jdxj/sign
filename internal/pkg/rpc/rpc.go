package rpc

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/logger"
)

const (
	registry = "/registry/service"
)

var (
	ErrInvalidServiceName = errors.New("invalid service name")
	ErrListenAddrNotFound = errors.New("listen addr not found")
)

func DSN(service string) string {
	// scheme:///service
	return fmt.Sprintf("%s:///%s", SignScheme, service)
}

func DSNLocal(service string) string {
	return fmt.Sprintf("%s:///%s", SignSchemeLocal, service)
}

func NewServer(port int) *Server {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	s := &Server{
		listener:   l,
		grpcServer: grpc.NewServer(),
		wg:         &sync.WaitGroup{},
	}
	return s
}

type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server

	wg *sync.WaitGroup
}

func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.grpcServer.RegisterService(desc, impl)
}

func (s *Server) Serve() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		logger.Infof("server started")
		err := s.grpcServer.Serve(s.listener)
		if err != nil {
			logger.Errorf("grpc serve err: %s", err)
		}
	}()
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
	s.wg.Wait()
	_ = s.listener.Close()
	logger.Infof("server stopped")
}

// Deprecated
func NewClient(target string, newClient func(cc *grpc.ClientConn)) {
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	newClient(cc)
}

func NewConn(service string) grpc.ClientConnInterface {
	cc, err := grpc.Dial(DSNLocal(service), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return cc
}

func GetListenAddr(port int, mod string) (string, error) {
	switch mod {
	default:
		return fmt.Sprintf("%s:%d", "127.0.0.1", port), nil
	case "release":
	}

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addresses {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		ip := ipNet.IP
		if !ip.IsLoopback() && ip.To4() != nil {
			listenAddr := fmt.Sprintf("%s:%d", ip.To4(), port)
			return listenAddr, nil
		}
	}
	return "", ErrListenAddrNotFound
}
