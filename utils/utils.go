package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	Pic58CookieURL    = "https://www.58pic.com"
	Pic58CookieDomain = ".58pic.com"
	//BilibiliCookieURL    = "https://www.bilibili.com"
	BilibiliCookieURL    = "https://space.bilibili.com"
	BilibiliCookieDomain = ".bilibili.com"
)

// StrToCookies 将给定的 cookie 字符串转换成 http.Cookie,
// domain 是 http.Cookie 所必须的.
func StrToCookies(cookiesStr, domain string) ([]*http.Cookie, error) {
	if domain == "" {
		return nil, fmt.Errorf("invaild domain")
	}

	cookiesStr = strings.ReplaceAll(cookiesStr, " ", "")
	cookiesParts := strings.Split(cookiesStr, "&")

	var cookies []*http.Cookie
	for _, part := range cookiesParts {
		idx := strings.Index(part, "=")
		if idx < 0 {
			MyLogger.Warn("%s not found '=' in cookie part: %s", Log_Log, part)
			continue
		}
		k := part[:idx]
		v := part[idx+1:]

		cookie := &http.Cookie{
			Name:     k,
			Value:    v,
			Path:     "/",
			Domain:   domain,
			Expires:  time.Now().Add(time.Hour * 24 * 365), // 一年后过期
			Secure:   false,
			HttpOnly: false,
		}
		cookies = append(cookies, cookie)
	}

	if len(cookies) == 0 {
		return nil, fmt.Errorf("invalid cookie")
	}
	return cookies, nil
}

func NowUnixMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
