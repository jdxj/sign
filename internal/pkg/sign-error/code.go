package sign_error

import (
	"net/http"
)

var (
	// codeMap 维护了 appCode 到 httpCode 的映射
	codeMap = make(map[int]Code)
	// validHTTP 有效的 httpCode
	validHTTP = map[int]struct{}{
		http.StatusOK:                  {},
		http.StatusMultipleChoices:     {},
		http.StatusBadRequest:          {},
		http.StatusInternalServerError: {},
	}
)

type Code struct {
	// http 状态码
	http int
	// 应用层状态码
	app int
	// 应用层状态码描述
	desc string
}

func (c Code) HTTP() int {
	if c.http == 0 {
		return http.StatusInternalServerError
	}
	return c.http
}

func (c Code) APP() int {
	return c.app
}

func (c Code) String() string {
	http.StatusOK
	return c.desc
}

func register(app, http int, desc string) {

}
