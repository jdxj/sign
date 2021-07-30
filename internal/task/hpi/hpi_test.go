package hpi

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/jdxj/sign/internal/pkg/bot"
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

	if err := SignIn(client, "jdxj"); err != nil {
		t.Fatalf("%s\n", "sign failed")
	}

}

func TestSignByStep(t *testing.T) {
	bot.Init(botKey, chatID)
	client, err := Auth(tmp)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	token, err := getSignToken(client)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("token: %s\n", token)

	err = accessSignURL(client, token)
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
