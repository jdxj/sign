package evo

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

type Updater struct{}

func (u *Updater) Kind() string {
	return pb.Kind_EVOLUTION_RELEASE.String()
}

func (u *Updater) Execute(body []byte) (string, error) {
	param := &pb.Evolution{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}
	url := fmt.Sprintf(buildURL, param.GetDevice())
	bi := &buildInfo{}
	err = util.ParseBody(&http.Client{}, url, bi)
	if err != nil {
		return msgEvoUpdateFailed, err
	}
	updateTime := time.Unix(bi.Datetime, 0)
	if time.Since(updateTime) <= 24*time.Hour {
		return bi.String(), nil
	}
	return "", nil
}
