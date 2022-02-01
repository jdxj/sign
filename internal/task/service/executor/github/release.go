package github

import (
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
)

const (
	msgReleaseUpdateFailed = "获取release更新失败"
	msgParseParamFailed    = "解析参数失败"
)

const (
	apiPrefix = "https://api.github.com"

	repositories = "/repos"
	// {owner}/{repo}
	releases = apiPrefix + repositories + "/%s/%s" + "/releases?per_page=1&page=1"
)

var (
	ErrReleaseNotFound = errors.New("release not found")
)

type Release struct{}

func (rel *Release) Kind() string {
	return pb.Kind_GITHUB_RELEASE.String()
}

func (rel *Release) Execute(body []byte) (string, error) {
	param := &pb.GithubRelease{}
	err := proto.Unmarshal(body, param)
	if err != nil {
		return msgParseParamFailed, err
	}

	rsp, err := getRelease(param.GetOwner(), param.GetRepo())
	if errors.Is(err, ErrReleaseNotFound) {
		return "", nil
	} else if err != nil {
		return msgReleaseUpdateFailed, err
	}

	if released(rsp) {
		return fmt.Sprintf("%s/%s 有新的 release", param.GetOwner(), param.GetRepo()), nil
	}
	return "", nil
}

type response struct {
	// UTC 时间
	CreatedAt string `json:"created_at"`
}

func getRelease(owner, repo string) (*response, error) {
	var rsps []*response
	url := fmt.Sprintf(releases, owner, repo)
	err := getJsonWithHeader(url, &rsps)
	if err != nil {
		return nil, err
	}
	if len(rsps) < 1 {
		return nil, ErrReleaseNotFound
	}
	return rsps[0], nil
}

func getJsonWithHeader(url string, rsp interface{}) error {
	header := map[string]string{
		"Accept": "application/vnd.github.v3+json",
	}
	return util.GetJson(url, header, rsp)
}

func released(rsp *response) bool {
	createdAt, _ := time.Parse(time.RFC3339, rsp.CreatedAt)
	if time.Since(createdAt) <= 24*time.Hour {
		return true
	}
	return false
}
