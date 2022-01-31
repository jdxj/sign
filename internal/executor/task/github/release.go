package github

import (
	"errors"
	"fmt"
	"time"

	"github.com/jdxj/sign/internal/pkg/util"
)

const (
	msgReleaseUpdateFailed = "获取release更新失败"
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

func (rel *Release) Domain() crontab.Domain {
	return crontab.Domain_Github
}

func (rel *Release) Kind() crontab.Kind {
	return crontab.Kind_Release
}

func (rel *Release) Execute(key string) (string, error) {
	req := &request{}
	err := util.PopulateStruct(key, req)
	if err != nil {
		return msgReleaseUpdateFailed, err
	}
	rsp, err := getRelease(req.Owner, req.Repo)
	if err != nil {
		return msgReleaseUpdateFailed, err
	}
	ok, err := released(rsp)
	if ok {
		return fmt.Sprintf("%s/%s 有新的 release", req.Owner, req.Repo), nil
	}
	return "", fmt.Errorf("release not found: %w", err)
}

type request struct {
	Owner string
	Repo  string
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

func released(rsp *response) (bool, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return false, err
	}
	createAt, err := time.ParseInLocation(time.RFC3339, rsp.CreatedAt, loc)
	if err != nil {
		return false, err
	}
	if time.Since(createAt) <= 24*time.Hour {
		return true, nil
	}
	return false, nil
}
