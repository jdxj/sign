package logger

import (
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
)

func TestMain(t *testing.M) {
	Init(config.Logger{
		Path: "./test.log",
	})
	os.Exit(t.Run())
}

func TestDebugf(t *testing.T) {
	defer Sync()

	Debugf("abc: %s", "haha")
	Infof("def: %s", "123")
	Errorf("ghi: %s", "456")
}
