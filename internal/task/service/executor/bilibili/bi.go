package bilibili

import (
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

type Bi struct{}

func (b *Bi) Kind() string {
	return pb.Kind_BILIBILI_B_COUNT.String()
}

func (b *Bi) Execute(body []byte) (string, error) {
	param := &pb.BiLiBiLi{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParam, err
	}
	c, err := auth(param.GetCookie())
	if err != nil {
		return msgGetBiFailed, err
	}
	return queryBi(c)
}

func queryBi(c *http.Client) (string, error) {
	biResp := &biResp{}
	err := util.ParseBody(c, biURL, biResp)
	if err != nil {
		return msgGetBiFailed,
			fmt.Errorf("stage: %s, error: %w", stageQuery, err)
	}

	if biResp.Code != 0 {
		return msgGetBiFailed,
			fmt.Errorf("%w, stage: %s", ErrInvalidCookie, stageQuery)
	}

	msg := fmt.Sprintf("硬币: %d", biResp.Data.Money)
	return msg, nil
}
