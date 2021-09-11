package rpc

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/logger"
)

const (
	registry = "/registry/service"
)

func DSN(service string) string {
	// scheme:///service
	return fmt.Sprintf("%s:///%s", SignScheme, service)
}

func NewServer(service, listenAddr, etcdAddr string) (*Server, error) {
	if service == "" {
		return nil, fmt.Errorf("invalid service name")
	}

	c, err := clientV3.New(clientV3.Config{
		Endpoints: []string{etcdAddr},
	})
	if err != nil {
		return nil, err
	}

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}

	s := &Server{
		etcdClient: c,
		listener:   l,
		grpcServer: grpc.NewServer(),
		listenAddr: listenAddr,
		node:       time.Now().UnixNano(),
		service:    service,
		stop:       make(chan int),
		wg:         &sync.WaitGroup{},
	}
	return s, nil
}

type Server struct {
	etcdClient *clientV3.Client
	listener   net.Listener
	grpcServer *grpc.Server

	listenAddr string
	node       int64
	service    string

	stop chan int
	wg   *sync.WaitGroup
}

func (s *Server) Register(registerServer func(grpc.ServiceRegistrar)) {
	registerServer(s.grpcServer)

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
				logger.Infof("stop keepalive")
				return
			case <-timer.C:
			}

			err := s.register()
			if err != nil {
				logger.Errorf("register: %s, listenAddr: %s err: %s",
					s.service, s.listenAddr, err)
			}
		}
	}()
}

func (s *Server) register() error {
	// todo: 配置化 TTL
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lease, err := s.etcdClient.Grant(ctx, 5)
	if err != nil {
		return err
	}
	_, err = s.etcdClient.Put(ctx, s.etcdKey(), s.listenAddr, clientV3.WithLease(lease.ID))
	if err == nil {
		logger.Debugf("register success, key: %s, value: %s",
			s.etcdKey(), s.listenAddr)
	}
	return err
}

func (s *Server) etcdKey() string {
	// prefix/serviceName/nodeID
	return fmt.Sprintf("%s/%s/%d",
		registry, s.service, s.node)
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
	close(s.stop)
	s.grpcServer.GracefulStop()
	s.wg.Wait()
	logger.Infof("server stopped")
}

func NewClient(service string, newClient func(conn *grpc.ClientConn)) {
	cc, err := grpc.Dial(DSN(service), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	newClient(cc)
}
