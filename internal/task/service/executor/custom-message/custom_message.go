package custom_message

import (
	"google.golang.org/protobuf/proto"

	pb "github.com/jdxj/sign/internal/proto/task"
)

const (
	msgParseParamFailed = "解析参数失败"
)

type CustomMessage struct{}

func (cm *CustomMessage) Kind() string {
	return pb.Kind_CUSTOM_MESSAGE.String()
}

func (cm *CustomMessage) Execute(body []byte) (string, error) {
	param := &pb.CustomMessage{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}
	return param.GetContent(), nil
}
