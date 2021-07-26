package hpi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/pkg/task/common"
)

const (
	Domain  = ".ld246.com"
	URL     = "https://ld246.com/"
	AuthURL = "https://ld246.com/notifications/unread/count?_=%d"
	SignURL = "https://ld246.com/activity/daily-checkin"
)

func Auth(cookies string) (*http.Client, error) {
	jar := common.NewJar(cookies, Domain, URL)
	client := &http.Client{Jar: jar}
	authResp := make(map[string]interface{})

	param := time.Now().UnixNano() / 1000000
	u := fmt.Sprintf(AuthURL, param)
	err := common.ParseBody(client, u, &authResp)
	if err != nil {
		return client, err
	}
	// 解析没问题应该就是成功了
	return client, nil
}

func SignIn(task *common.Task) bool {
	signResp := make(map[string]interface{})
	err := common.ParseBody(task.Client, SignURL, &signResp)
	if err != nil {
		text := fmt.Sprintf("签到失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}
	fmt.Printf("%#v\n", signResp)
	return false
}

func verify(client *http.Client) error {
	return nil
}
