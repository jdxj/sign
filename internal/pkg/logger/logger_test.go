package logger

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	Init("./test_logger.log")
	os.Exit(t.Run())
}

func TestDebugf(t *testing.T) {
	Debugf("abc: %s", "haha")
}
