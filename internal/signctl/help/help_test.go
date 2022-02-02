package help

import (
	"fmt"
	"testing"

	"github.com/jdxj/sign/internal/proto/task"
)

func TestGenParamList(t *testing.T) {
	res := getParamList(task.Kind_STG_SIGN_IN.String())
	fmt.Printf("res: %s\n", res)
}
