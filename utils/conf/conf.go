package conf

import "gopkg.in/ini.v1"

var (
	Conf *ini.File
)

func init() {
	conf, err := ini.Load("sign.ini")
	if err != nil {
		panic(err)
	}

	Conf = conf
}
