package hpi

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/pkg/task/common"
)

const (
	Domain       = ".ld246.com"
	URL          = "https://ld246.com/"
	AuthURL      = "https://ld246.com/notifications/unread/count?_=%d"
	SignTokenURL = "https://ld246.com/activity/checkin"
	SignURL      = "https://ld246.com/activity/daily-checkin?token=%s"
	VerifyURL    = "https://ld246.com/member/jdxj/points?p=1&pjax=true"
)

var (
	regSignToken *regexp.Regexp
	regVerify    *regexp.Regexp
)

func init() {
	regSignToken = regexp.MustCompile(`csrfToken: '(.+)'`)
	regVerify = regexp.MustCompile(`202\d-\d{2}-\d{2}`)
}

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

// hpi 签到步骤：
//   1. 获取 sign token
//   2. 访问 sign url
//   3. 验证

func SignIn(task *common.Task) bool {
	st, err := getSignToken(task.Client)
	if err != nil {
		text := fmt.Sprintf("签到失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}

	err = accessSignURL(task.Client, st)
	if err != nil {
		text := fmt.Sprintf("签到失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}

	err = verify(task.Client)
	if err != nil {
		text := fmt.Sprintf("签到失败, id: %s, type: %s, err: %s",
			task.ID, common.TypeMap[task.Type], err)
		bot.Send(text)
		return false
	}

	text := fmt.Sprintf("签到成功, id: %s, type: %s",
		task.ID, common.TypeMap[task.Type])
	bot.Send(text)
	return true
}

func getSignToken(client *http.Client) (string, error) {
	body, err := common.ParseRawBody(client, SignTokenURL)
	if err != nil {
		return "", err
	}

	matched := regSignToken.FindStringSubmatch(string(body))
	if len(matched) != 2 {
		return "", fmt.Errorf("sign token not found: %v", matched)
	}
	return matched[1], nil
}

func accessSignURL(client *http.Client, token string) error {
	u := fmt.Sprintf(SignURL, token)
	return common.ParseBody(client, u, nil)
}

func verify(client *http.Client) error {
	d, err := common.ParseRawBody(client, VerifyURL)
	if err != nil {
		return err
	}
	date := regVerify.FindString(string(d))
	return common.VerifyDate(date)
}
