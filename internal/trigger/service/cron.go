package service

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

const (
	// randomly is used to generate random time.
	// example:
	//   '@randomly'                  means 00:00:00~23:59:59
	//   '@randomly 8:00:00'          means 8:00:00~23:59:59
	//   '@randomly 8:00:00 23:00:00' means 8:00:00~23:00:00
	randomly     = "@randomly"
	randomLayout = "15:04:05"
)

var (
	ErrInvalidRandomly = errors.New("invalid randomly expression")
)

func newRandomSchedule(spec string) (*randomSchedule, error) {
	spec = strings.TrimSpace(strings.TrimPrefix(spec, randomly))

	var (
		ranges = strings.Split(spec, " ")
		start  time.Time
		end    time.Time
		err    error
	)

	switch len(ranges) {
	case 1:
		if ranges[0] == "" {
			start = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
		}
		if start, err = time.ParseInLocation(randomLayout, ranges[0], time.Local); err != nil {
			return nil, err
		}
		end = time.Date(0, 0, 0, 23, 59, 59, 0, time.Local)
	case 2:
		start, err = time.ParseInLocation(randomLayout, ranges[0], time.Local)
		if err != nil {
			return nil, err
		}
		end, err = time.ParseInLocation(randomLayout, ranges[1], time.Local)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrInvalidRandomly
	}

	if start.After(end) {
		return nil, ErrInvalidRandomly
	}

	return &randomSchedule{
		start: start,
		end:   end,
	}, nil
}

type randomSchedule struct {
	start time.Time
	end   time.Time
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
	rs.end.Sub(rs.start).Truncate()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).
		AddDate(0, 0, 1).
		Add(time.Duration(offset) * time.Second)
}

func (s *Service) parse(spec string) (cron.Schedule, error) {
	switch {
	case strings.HasPrefix(spec, randomly):
		return newRandomSchedule(spec)
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
