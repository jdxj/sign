package utils

import (
	"math/rand"
	"net/http"
	"time"
)

// DailyRandTimeExec 用于每日签到, 做了随机
func DailyRandTimeExec(prefix string, f func()) {
	timer := time.NewTimer(time.Hour)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 无限循环
	for {
		// 现在时刻
		now := time.Now()
		// 明天这个时候
		tomNow := now.Add(24 * time.Hour)
		// 明天0点
		tomZero := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
		// 明天随便几点
		inc := time.Duration(r.Intn(24 * 60 * 60))
		tomSome := tomZero.Add(inc * time.Second)
		// 下次签到时延
		dur := tomSome.Sub(now)

		timer.Reset(dur)
		MyLogger.Debug("%s 等待时间到达: %s", prefix, tomSome)

		select {
		case <-timer.C:
			f()
		}

		MyLogger.Debug("%s %s", prefix, "本次每日任务完成...")
	}
}

const (
	Cookie_Bilibili = ".bilibili.com"
	Cookie_58pic    = ".58pic.com"
)

func Cookies(prefix, domain string) []*http.Cookie {
	var cookies []*http.Cookie

	kvs := ConfAll(prefix)
	// 无所谓的过期时间
	expires := time.Date(2048, 1, 1, 0, 0, 0, 0, time.Now().Location())
	for _, kv := range kvs {
		cookie := &http.Cookie{
			Name:     kv.K,
			Value:    kv.V,
			Path:     "/",
			Domain:   domain,
			Expires:  expires,
			Secure:   false,
			HttpOnly: false,
		}

		cookies = append(cookies, cookie)
	}

	if len(cookies) != 0 {
		MyLogger.Debug("%s %s", prefix, "读取配置成功")
	}
	return cookies
}
