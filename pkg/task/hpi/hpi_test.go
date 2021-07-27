package hpi

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/pkg/task/common"
)

func TestMSecond(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(1627303261118)
}

var (
	tmp          = ""
	botKey       = ""
	chatID int64 = 0
)

func TestAuth(t *testing.T) {
	bot.Init(botKey, chatID)
	client, err := Auth(tmp)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	task := &common.Task{
		ID:     "test",
		Type:   202,
		Client: client,
	}
	if !SignIn(task) {
		t.Fatalf("%s\n", "sign failed")
	}

}

func TestSignTokenURLBody(t *testing.T) {
	client, err := Auth(tmp)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	st, err := getSignToken(client)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	err = accessSignURL(client, st)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	err = verify(client)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestFindSignToken(t *testing.T) {
	str := ``
	reg := regexp.MustCompile(`csrfToken: '(.+)'`)
	res := reg.FindStringSubmatch(str)
	fmt.Println(res)
}

func TestRegDate(t *testing.T) {
	str := ``
	reg := regexp.MustCompile(str)
	res := reg.FindAllString(str, -1)
	for _, d := range res {
		fmt.Println(d)
	}
}
