package rpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/config"
)

const (
	registry = "/registry/service"
)

var (
	ErrInvalidServiceName = errors.New("invalid service name")
	ErrListenAddrNotFound = errors.New("listen addr not found")
)

type option struct {
	debug bool
}

type OptFunc func(*option)

func WithDebug() OptFunc {
	return func(opt *option) {
		opt.debug = true
	}
}

func DSN(service string) string {
	if !isLocal {
		return service
	}
	return fmt.Sprintf("%s:///%s", SignScheme, service)
}

func NewServer(service, node string, cfg config.RPC, optFs ...OptFunc) *Server {
	opt := new(option)
	for _, f := range optFs {
		f(opt)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	s := &Server{
		listener:   l,
		grpcServer: grpc.NewServer(),
		service:    service,
		node:       node,
		opt:        opt,
		addr:       addr,
		wg:         &sync.WaitGroup{},
		stop:       make(chan int),
	}

	if !opt.debug {
		s.etcdClient = NewEtcdClient(cfg.Endpoints, cfg.Ca, cfg.Cert, cfg.Key)

		s.wg.Add(1)
		go func() {
			defer s.wg.Done()

			dur := 5 * time.Second
			timer := time.NewTimer(dur)
			defer timer.Stop()

			for {
				timer.Reset(dur)
				select {
				case <-s.stop:
					log.Printf("stop register address: %s", s.ID())
					return
				case <-timer.C:
					s.registerAddress()
				}
			}
		}()
	}
	return s
}

type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	etcdClient *etcd.Client

	service string
	node    string
	opt     *option
	addr    string

	wg   *sync.WaitGroup
	stop chan int
}

// RegisterService 不要直接调用该方法, 应使用 pb 中的 RegisterXxxServer()
func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.grpcServer.RegisterService(desc, impl)
}

func (s *Server) registerAddress() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lease, err := s.etcdClient.Grant(ctx, 5)
	if err != nil {
		log.Printf("etcd grant err: %s, id: %s\n", err, s.ID())
		return
	}
	_, err = s.etcdClient.Put(ctx, s.ID(), s.addr, etcd.WithLease(lease.ID))
	if err != nil {
		log.Printf("etcd put err: %s, id: %s\n", err, s.ID())
	}
}

func (s *Server) Serve() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		log.Println("server started")
		err := s.grpcServer.Serve(s.listener)
		if err != nil {
			log.Printf("grpc serve err: %s\n", err)
		}
	}()
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
	s.wg.Wait()
	_ = s.listener.Close()
	log.Println("server stopped")
}

func (s *Server) ID() string {
	return fmt.Sprintf("%s/%s/%s", registry, s.service, s.node)
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
	cc, err := grpc.Dial(DSN(service), grpc.WithInsecure())
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
