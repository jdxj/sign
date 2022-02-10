package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNew(t *testing.T) {
	r := gin.New()
	New(r)

	err := r.Run(":8080")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
