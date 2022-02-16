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

	randomStart = "00:00:00"
	randomEnd   = "23:59:59"
)

var (
	ErrInvalidRandomly = errors.New("invalid randomly expression")

	randomStartTime, _ = time.ParseInLocation(randomLayout, randomStart, time.UTC)
)

func parseRanges(ts []string) ([]time.Time, error) {
	result := make([]time.Time, 0, len(ts))
	for _, v := range ts {
		if v == "" {
			result = append(result, randomStartTime)
			continue
		}

		t, err := time.ParseInLocation(randomLayout, v, time.UTC)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func newRandomSchedule(spec string) (*randomSchedule, error) {
	spec = strings.TrimSpace(strings.TrimPrefix(spec, randomly))
	ranges := strings.Split(spec, " ")

	if len(ranges) == 1 {
		ranges = append(ranges, randomEnd)
	} else if len(ranges) > 2 {
		return nil, ErrInvalidRandomly
	}
	result, err := parseRanges(ranges)
	if err != nil {
		return nil, err
	}

	start, end := result[0], result[1]
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

func (rs *randomSchedule) offset() time.Duration {
	d1 := rs.end.Sub(rs.start)
	if d1.Nanoseconds() == 0 {
		return 0
	}
	d2 := time.Duration(rand.Int63n(d1.Nanoseconds()))
	d3 := rs.start.Sub(randomStartTime)
	return d2 + d3
}

// Next 返回明天的某个时间
func (rs *randomSchedule) Next(t time.Time) time.Time {
	var (
		year  = t.Year()
		month = t.Month()
		day   = t.Day()
	)
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).
		AddDate(0, 0, 1).
		Add(rs.offset())
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
