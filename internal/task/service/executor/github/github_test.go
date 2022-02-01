package github

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/proto/task"
)

func TestRelease_Execute(t *testing.T) {
	exe := &Release{}
	param := &task.GithubRelease{
		Owner: "asim",
		Repo:  "go-micro",
	}
	d, err := proto.Marshal(param)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	str, err := exe.Execute(d)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", str)
}

func TestSince(t *testing.T) {
	tt := time.Now().Add(time.Hour)
	fmt.Printf("%v\n", time.Since(tt))
}

func TestReleased(t *testing.T) {
	// mock UTC
	loc, err := time.LoadLocation("Africa/Abidjan")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	now := time.Now().In(loc)
	cases := []struct {
		Name   string
		Rsp    *response
		Expect bool
	}{
		{
			Name: "case1",
			Rsp: &response{
				CreatedAt: now.Add(-25 * time.Hour).Format(time.RFC3339),
			},
			Expect: false, // 无更新
		},
		{
			Name: "case2",
			Rsp: &response{
				CreatedAt: now.Add(-10 * time.Hour).Format(time.RFC3339),
			},
			Expect: true, // 有更新
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := released(c.Rsp)
			if res != c.Expect {
				t.Errorf("expect: %t, get: %t\n", c.Expect, res)
			}
		})
	}
}

func TestErrorsIs(t *testing.T) {
	err := ErrReleaseNotFound

	if !errors.Is(err, ErrReleaseNotFound) {
		t.Fatalf("failed")
	}
}
