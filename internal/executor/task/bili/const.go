package bili

import (
	"errors"
)

const (
	domain    = ".bilibili.com"
	signURL   = "https://www.bilibili.com/"
	authURL   = "https://api.bilibili.com/x/member/web/account"
	verifyURL = "https://api.bilibili.com/x/member/web/coin/log?jsonp=jsonp"
	biURL     = "https://account.bilibili.com/site/getCoin"
)

const (
	msgSignInFailed = "B站签到失败"
	msgGetBiFailed  = "获取B币失败"
)

var (
	ErrLogNotFound   = errors.New("log not found")
	ErrSignIn        = errors.New("sign in failed")
	ErrInvalidCookie = errors.New("invalid cookie")
)

type authResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		MID    int    `json:"mid"`
		Uname  string `json:"uname"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

type biResp struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
	Data   struct {
		Money int `json:"money"`
	} `json:"data"`
}

type verifyResp struct {
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
