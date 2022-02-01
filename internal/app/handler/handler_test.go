package handler

import (
	"encoding/json"
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/proto/task"
)

func TestEncodeCreateTask(t *testing.T) {
	param := &task.CustomMessage{Content: "hello world"}
	body, err := proto.Marshal(param)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	ctReq := &CreateTaskReq{
		Desc:  "",
		Kind:  "",
		Spec:  "",
		Param: body,
	}
	d, err := json.Marshal(ctReq)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("d: %s\n", d)
}
