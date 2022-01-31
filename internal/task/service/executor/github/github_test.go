package github

import (
	"fmt"
	"testing"
	"time"
)

func TestRelease_Execute(t *testing.T) {
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
			res, err := released(c.Rsp)
			if err != nil {
				t.Fatalf("%s\n", err)
			}
			if res != c.Expect {
				t.Errorf("expect: %t, get: %t\n", c.Expect, res)
			}
		})
	}
}

func TestReleased2(t *testing.T) {
	rsp := &response{CreatedAt: "2022-01-31T11:46:00Z"}
	_, err := released(rsp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
