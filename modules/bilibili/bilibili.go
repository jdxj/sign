package bilibili

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	"sign/utils/conf"
	"sign/utils/log"
	"strconv"
)

func NewBiliFromApi(conf *conf.BiliConf) (*ToucherBilibili, error) {
	if conf == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	following := strconv.Itoa(conf.VerifyValue)

	t := &ToucherBilibili{
		name:        conf.Name,
		cookies:     conf.Cookies,
		loginURL:    "https://api.bilibili.com/x/web-interface/nav/stat",
		verifyValue: following,
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil

}
func NewToucherBilibili(sec *ini.Section) (*ToucherBilibili, error) {
	if sec == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &ToucherBilibili{
		name:        sec.Name(),
		cookies:     sec.Key("cookies").String(),
		loginURL:    "https://api.bilibili.com/x/web-interface/nav/stat",
		verifyValue: sec.Key("verifyValue").String(),
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

type ToucherBilibili struct {
	name        string
	cookies     string
	loginURL    string
	verifyValue string

	client *http.Client
}

func (tou *ToucherBilibili) Name() string {
	return tou.name
}

func (tou *ToucherBilibili) Boot() bool {
	cookies, err := utils.StrToCookies(tou.cookies, utils.BilibiliCookieDomain)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	cookieURL, err := url.Parse(utils.BilibiliCookieURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

func (tou *ToucherBilibili) Login() bool {
	req, err := http.NewRequest("GET", tou.loginURL, nil)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	// todo: 为所有请求生成 user-agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	resp, err := tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", err)
		return false
	}

	var loginResp loginResp
	if err := json.Unmarshal(data, &loginResp); err != nil {
		log.MyLogger.Error("%s %s", err)
		return false
	}

	if strconv.Itoa(loginResp.Data.Following) == tou.verifyValue {
		return true
	}
	return false
}

func (tou *ToucherBilibili) Sign() bool {
	_, err := tou.client.Get("http://bilibili.com")
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_Bilibili, err)
		return false
	}

	return true
}

type loginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Following    int `json:"following"`
		Follower     int `json:"follower"`
		DynamicCount int `json:"dynamic_count"`
	} `json:"data"`
}
