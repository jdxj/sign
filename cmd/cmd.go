package cmd

import (
	"fmt"
	"gopkg.in/ini.v1"
	"sign/modules"
	pic "sign/modules/58pic"
	"sign/modules/bilibili"
	"sign/modules/studygolang"
)

type Site int

const (
	Pic58 Site = iota + 1
	StudyGolang
	Bilibili
)

// todo 获取配置文件
func NewToucher(sec *ini.Section) (modules.Toucher, error) {
	if sec == nil {
		return nil, fmt.Errorf("invaild cfg")
	}

	site, err := sec.Key("site").Int()
	if err != nil {
		return nil, err
	}

	switch Site(site) {
	case Pic58:
		return pic.NewToucher58Pic(sec)
	case StudyGolang:
		return bilibili.NewToucherBilibili(sec)
	case Bilibili:
		return studygolang.NewToucherStudyGolang(sec)
	}

	return nil, fmt.Errorf("did not implement this site")
}
