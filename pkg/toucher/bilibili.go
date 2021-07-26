package toucher

import (
	"net/http"
	"time"
)

type AuthRespBili struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		MID    int    `json:"mid"`
		Uname  string `json:"uname"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

type VerifyRespBili struct {
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

func NewBili(id, cookies string) (*Bili, error) {
	bili := &Bili{
		id:     id,
		client: &http.Client{},
	}
	return bili, bili.Auth(cookies)
}

type Bili struct {
	id     string
	client *http.Client
}

func (bili *Bili) ID() string {
	return bili.id
}

func (bili *Bili) Domain() string {
	return DomainBili
}

func (bili *Bili) Auth(cookies string) error {
	jar := NewJar(cookies, DomainBili, urlBili)
	bili.client.Jar = jar

	authResp := &AuthRespBili{}
	err := ParseBody(bili.client, authBili, authResp)
	if err != nil {
		return err
	}

	if authResp.Code != 0 {
		return ErrorAuthFailed
	}
	return nil
}

func (bili *Bili) SignIn() error {
	return ParseBody(bili.client, urlBili, nil)
}

func (bili *Bili) Verify() error {
	verifyResp := &VerifyRespBili{}
	err := ParseBody(bili.client, verifyBili, verifyResp)
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
