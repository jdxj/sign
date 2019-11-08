package hacpai

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"sign/utils/log"
)

func NewToucherHacPai(sec *ini.Section) (*ToucherHacPai, error) {
	if sec == nil {
		return nil, fmt.Errorf("invaild section config")
	}

	tou := &ToucherHacPai{
		username:   sec.Key("username").String(),
		password:   sec.Key("password").String(),
		loginURL:   "https://hacpai.com/api/v2/login",
		signRefURL: "https://hacpai.com/activity/checkin",
		signURL:    "https://hacpai.com/activity/daily-checkin",
		client:     &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tou.client.Jar = jar
	return tou, nil
}

type ToucherHacPai struct {
	username string
	password string // 需要 md5 值
	token    string

	loginURL   string
	signRefURL string
	signURL    string

	client *http.Client
}

func (tou *ToucherHacPai) Name() string {
	return ""
}

func (tou *ToucherHacPai) Boot() bool {
	return true
}

func (tou *ToucherHacPai) Login() bool {
	loginData := make(map[string]interface{})
	loginData["userName"] = tou.username
	loginData["userPassword"] = toMd5(tou.password)
	loginData["captcha"] = ""
	loginDataJson, err := json.Marshal(loginData)
	if err != nil {
		log.MyLogger.Error("%s marshal login data err: %s", log.Log_HacPai, err)
		return false
	}
	reader := bytes.NewReader(loginDataJson)
	req, err := http.NewRequest("POST", tou.loginURL, reader)
	if err != nil {
		log.MyLogger.Error("%s new login req err: %s", log.Log_HacPai, err)
		return false
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")
	resp, err := tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s err when send login data: %s", log.Log_HacPai, err)
		return false
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s err when read login resp: %s", log.Log_HacPai, err)
		return false
	}

	respDate := make(map[string]interface{})
	if err = json.Unmarshal(respBytes, &respDate); err != nil {
		log.MyLogger.Error("%s err when unmarshal login resp: %s", log.Log_HacPai, err)
		return false
	}

	if tmp, ok := respDate["token"]; ok {
		if token, mark := tmp.(string); mark && token != "" {
			tou.token = token
			return mark
		}
	}

	log.MyLogger.Debug("%s login token not found", log.Log_HacPai)
	return false
}

func (tou *ToucherHacPai) Sign() bool {
	req, err := http.NewRequest("GET", tou.signURL, nil)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")
	req.Header.Set("Referer", tou.signRefURL)
	cookie := http.Cookie{
		Name:   "symphony",
		Value:  tou.token,
		Path:   "/",
		MaxAge: 86400,
	}
	req.AddCookie(&cookie)

	resp, err := tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s err when send sign req: %s", log.Log_HacPai, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_HacPai, err)
		return false
	}

	target := doc.Find(".module__body").Find(".btn")
	if target == nil {
		log.MyLogger.Error("%s score not found, web page maybe change", log.Log_HacPai)
		return false
	} else {
		log.MyLogger.Info("%s score is: %s", log.Log_HacPai, target.Text())
		return true
	}
}

func toMd5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
