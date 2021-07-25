package toucher

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

var (
	BiliTch = NewBilibili()
)

type AuthRespBilibili struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		MID    int    `json:"mid"`
		Uname  string `json:"uname"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

type VerifyRespBilibili struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List []struct {
			Time   string `json:"time"`
			Delta  int    `json:"delta"`
			Reason string `json:"reason"`
		} `json:"list"`
		Count int `json:"count"`
	} `json:"data"`
}

func NewBilibili() *Bilibili {
	bili := &Bilibili{
		client: &http.Client{},
	}
	return bili
}

type Bilibili struct {
	client *http.Client
}

func (bili *Bilibili) RemoveJar() {
	bili.client.Jar = nil
}

func (bili *Bilibili) Domain() string {
	return DomainBili
}

func (bili *Bilibili) Auth(key string) (*cookiejar.Jar, error) {
	c := bili.client

	jar := NewJar(key, DomainBili, urlBili)
	c.Jar = jar
	defer bili.RemoveJar()

	authResp := &AuthRespBilibili{}
	err := ParseBody(c, authBili, authResp)
	if err != nil {
		return nil, err
	}

	if authResp.Code != 0 {
		return nil, ErrorAuthFailed
	}
	return jar, nil
}

func (bili *Bilibili) SignIn(jar *cookiejar.Jar) error {
	c := bili.client
	c.Jar = jar
	defer bili.RemoveJar()

	return ParseBody(c, urlBili, &Empty{})
}

func (bili *Bilibili) Verify(jar *cookiejar.Jar) error {
	c := bili.client
	c.Jar = jar
	defer bili.RemoveJar()

	verifyResp := &VerifyRespBilibili{}
	err := ParseBody(c, verifyBili, verifyResp)
	if err != nil {
		return err
	}

	if verifyResp.Code != 0 {
		return ErrorAuthFailed
	}
	list := verifyResp.Data.List
	if len(list) <= 0 {
		return ErrorLogNotFound
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	now := time.Now().In(loc)
	last, err := time.ParseInLocation("2006-01-02 15:04:05", list[0].Time, loc)
	if err != nil {
		return err
	}

	if now.YearDay() != last.YearDay() {
		return ErrorSignInFailed
	}
	return nil
}
