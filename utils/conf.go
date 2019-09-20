package utils

import "gopkg.in/ini.v1"

const (
	Conf_StudyGolang = "studygolang.com"
)

func Conf(prefix string, keys ...string) (res []string) {
	cfg, err := ini.Load("sign.ini")
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		value := cfg.Section(prefix).Key(key).String()
		res = append(res, value)
	}

	return
}
