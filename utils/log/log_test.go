package log

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("test")
}

func TestMyLog(t *testing.T) {

}

func TestMyLogger(t *testing.T) {
	MyLogger.Info("%s %s", Log_58pic, "test")
}

func TestPrintMap(t *testing.T) {
	m := make(map[string]string)
	m["abc"] = "cba"
	m["bcd"] = "dcb"
	m["cde"] = "edc"

	fmt.Printf("%s", m)
}
