package apiserver

import (
	"testing"
	"time"

	"github.com/jdxj/sign/internal/pkg/config"
)

func TestStart(t *testing.T) {
	Start(config.APIServer{Port: "8080"})
	time.Sleep(time.Hour)
}
