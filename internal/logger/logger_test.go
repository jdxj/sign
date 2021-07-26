package logger

import (
	"testing"

	"github.com/jdxj/sign/pkg/config"
)

func TestDebugf(t *testing.T) {
	root := config.ReadConfigs("/home/jdxj/workspace/sign/pkg/config/config.yaml")
	Init(root.Logger.Path, WithMode("release"))

	Debugf("%s, ac", "hah")
	Infof("hah")
}
