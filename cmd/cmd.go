package cmd

import (
	"fmt"
	"gopkg.in/ini.v1"
	"sign/modules"
	pic "sign/modules/58pic"
	"sign/modules/bilibili"
	"sign/modules/hacpai"
	"sign/modules/studygolang"
	"sign/modules/v2ex"
	"sign/utils/log"
)

type Site int

const (
	Pic58 Site = iota + 1
	StudyGolang
	Bilibili
	HacPai
	V2ex
)

func NewToucher(sec *ini.Section) (modules.Toucher, error) {
	if sec == nil {
		return nil, fmt.Errorf("invaild section config")
	}
	defer log.MyLogger.Info("%s load %s section config", log.Log_Cmd, sec.Name())

	site, err := sec.Key("site").Int()
	if err != nil {
		return nil, err
	}

	switch Site(site) {
	case Pic58:
		return pic.NewToucher58Pic(sec)
	case StudyGolang:
		return studygolang.NewToucherStudyGolang(sec)
	case Bilibili:
		return bilibili.NewToucherBilibili(sec)
	case HacPai:
		return hacpai.NewToucherHacPai(sec)
	case V2ex:
		return v2ex.NewToucherV2ex(sec)
	}

	return nil, fmt.Errorf("did not implement this site")
}
