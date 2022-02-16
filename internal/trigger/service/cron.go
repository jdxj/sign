package service

import (
	"math/rand"
	"time"

	"github.com/robfig/cron/v3"
)

const (
	randomly = "@randomly"
)

type randomSchedule struct {
}

// Next 返回明天的某个时间
func (rs *randomSchedule) Next(t time.Time) time.Time {
	var (
		year  = t.Year()
		month = t.Month()
		day   = t.Day()

		// 留5分钟执行任务, 避免延迟导致当天未执行.
		bound  = 24*60*60 - 5*60
		offset = rand.Intn(bound)
	)
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).
		AddDate(0, 0, 1).
		Add(time.Duration(offset) * time.Second)
}

func (s *Service) parse(spec string) (cron.Schedule, error) {
	switch spec {
	case randomly:
		return &randomSchedule{}, nil
	}
	return s.parser.Parse(spec)
}

func (s *Service) addJob(spec string, cmd cron.Job) error {
	schedule, err := s.parse(spec)
	if err != nil {
		return err
	}
	s.cron.Schedule(schedule, cmd)
	return nil
}
