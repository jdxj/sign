package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jdxj/sign/internal/pkg/logger"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

func NewServer(host string, port int, router http.Handler) *Server {
	if host == "" || port == 0 {
		panic(ErrInvalidConfig)
	}

	s := &Server{
		wg: &sync.WaitGroup{},
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", host, port),
			Handler: router,
		},
	}
	return s
}

type Server struct {
	wg         *sync.WaitGroup
	httpServer *http.Server
}

func (s *Server) Start() {
	s.wg.Add(1)
	go func() {
		s.wg.Done()

		logger.Infof("api server started")
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("ListenAndServe err: %s", err)
		}
	}()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = s.httpServer.Shutdown(ctx)
	s.wg.Wait()
	logger.Infof("api server already stop")
}
