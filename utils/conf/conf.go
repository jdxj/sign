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

const ExpectedLimit = 20

// api 用
type SGConf struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	ActiveURL string `json:"activeURL"`
	Expected  int    `json:"expected"`
	To        string `json:"to"`
}

func (sgc *SGConf) CheckValidity() bool {
	if sgc.Name == "" || sgc.Username == "" || sgc.Password == "" ||
		sgc.ActiveURL == "" || sgc.To == "" {

		return false
	}

	// 修正非法值
	if sgc.Expected < 0 || sgc.Expected > ExpectedLimit {
		sgc.Expected = 0
	}
	return true
}

type BiliConf struct {
	Name        string `json:"name"`
	Cookies     string `json:"cookies"`
	VerifyValue int    `json:"verify_value"`
	To          string `json:"to"`
}

func (bl *BiliConf) CheckValidity() bool {
	if bl.Name == "" || bl.Cookies == "" || bl.To == "" {
		return false
	}

	return true
}

type Pic58Conf struct {
	Name    string `json:"name"`
	Cookies string `json:"cookies"`
	To      string `json:"to"`
}

func (pic *Pic58Conf) CheckValidity() bool {
	if pic.Name == "" || pic.Cookies == "" || pic.To == "" {
		return false
	}

	return true
}

type HacPaiConf struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	To       string `json:"to"`
}

func (hac *HacPaiConf) CheckValidity() bool {
	if hac.Name == "" || hac.Username == "" || hac.Password == "" ||
		hac.To == "" {
		return false
	}

	return true
}

type V2exConf struct {
	Name    string `json:"name"`
	Cookies string `json:"cookies"`
	To      string `json:"to"`
}

func (v2ex *V2exConf) CheckValidity() bool {
	if v2ex.Name == "" || v2ex.Cookies == "" ||
		v2ex.To == "" {
		return false
	}

	return true
}

type IQiYiConf struct {
	Name        string `json:"name"`
	Cookies     string `json:"cookies"`
	CheckInSign string `json:"check_in_sign"`
	HotSpotSign string `json:"hot_spot_sign"`
	To          string `json:"to"`
}

func (iqy *IQiYiConf) CheckValidity() bool {
	if iqy.Name == "" || iqy.Cookies == "" || iqy.To == "" ||
		iqy.CheckInSign == "" || iqy.HotSpotSign == "" {
		return false
	}

	return true
}
