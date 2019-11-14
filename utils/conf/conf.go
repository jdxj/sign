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

// api 用
type SGConf struct {
	Site      string `json:"site"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	ActiveURL string `json:"activeURL"`
}

func (sgc *SGConf) CheckValidity() bool {
	if sgc.Site == "" || sgc.Username == "" || sgc.Password == "" || sgc.ActiveURL == "" {
		return false
	}

	return true
}
