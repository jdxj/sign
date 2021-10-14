package juejin

import (
	"github.com/jdxj/sign/internal/proto/crontab"
)

type Calendar struct{}

func (cal *Calendar) Domain() crontab.Domain {
	return crontab.Domain_JueJin
}

func (cal *Calendar) Kind() crontab.Kind {
	return crontab.Kind_JueJinCalendar
}

func (cal *Calendar) Execute(key string) (string, error) {
	return execute(key, calendarURL, &jokes{})
}
