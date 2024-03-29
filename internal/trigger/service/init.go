package service

import (
	"context"
	"sync"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/trigger/dao"
)

func New() *Service {
	s := &Service{
		mutex:  &sync.RWMutex{},
		specs:  make(map[string]struct{}),
		parser: cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor),
	}
	s.cron = cron.New(cron.WithParser(s.parser))
	s.start()
	return s
}

type Service struct {
	// 防止重复建立定时器
	mutex *sync.RWMutex
	specs map[string]struct{}

	cron   *cron.Cron
	parser cron.Parser
}

func (s *Service) start() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rows []dao.Specification
	err := db.WithCtx(ctx).
		Find(&rows).
		Error
	if err != nil {
		panic(err)
	}

	for _, v := range rows {
		s.specs[v.Spec] = struct{}{}
		err = s.addJob(v.Spec, newJob(v.Spec))
		if err != nil {
			logger.Errorf("AddFunc: %s, specID: %d\n", err, v.SpecID)
			continue
		}
		logger.Debugf("AddJob-specID: %d", v.SpecID)
	}
	s.cron.Start()
}

func (s *Service) hasAndAdd(spec string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.specs[spec]
	if exists {
		return true
	}
	s.specs[spec] = struct{}{}
	return false
}

func (s *Service) Stop() {
	<-s.cron.Stop().Done()
}
