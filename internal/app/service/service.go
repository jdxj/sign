package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/app/router"
	"github.com/jdxj/sign/internal/pkg/config"
	pb "github.com/jdxj/sign/internal/proto/app"
)

const (
	// todo: 自动化更新?
	version = "0.1.0"
)

func New(conf config.App) *Service {
	s := &Service{
		wg: &sync.WaitGroup{},
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", conf.Host, conf.Port),
			Handler: router.New(),
		},
	}
	s.start()
	return s
}

type Service struct {
	wg         *sync.WaitGroup
	httpServer *http.Server
}

func (s *Service) start() {
	s.wg.Add(1)
	go func() {
		s.wg.Done()

		log.Printf(" start app\n")
		err := s.httpServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			return
		} else if err != nil {
			log.Printf("ListenAndServe: %s", err)
		}
	}()
}

func (s *Service) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = s.httpServer.Shutdown(ctx)
	s.wg.Wait()
	log.Printf(" stop app")
}

func (s *Service) Version(_ context.Context, _ *emptypb.Empty, rsp *pb.VersionResponse) error {
	rsp.Version = version
	return nil
}
