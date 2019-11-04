package cmd

import (
	"fmt"
	"sign/modules"
	pic "sign/modules/58pic"
)

type App int

const (
	Pic58 App = iota
	StudyGolang
	Bilibili
)

// todo 获取配置文件
func NewToucher(app App) (modules.Toucher, error) {
	switch app {
	case Pic58:
		return pic.New58Pic(nil)
	case StudyGolang:
	case Bilibili:
	}

	return nil, fmt.Errorf("did not implement this type")
}
