package sign_error

import (
	"errors"
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

var (
	ErrInvalidHTTPCode          = errors.New("invalid http code")
	ErrAlreadyRegisteredAPPCode = errors.New("already registered app code")
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
	return c.desc
}

func register(http, app int, desc string) {
	_, exists := validHTTP[http]
	if exists {
		panic(ErrInvalidHTTPCode)
	}
	_, exists = codeMap[app]
	if exists {
		panic(ErrAlreadyRegisteredAPPCode)
	}

	codeMap[app] = Code{
		http: http,
		app:  app,
		desc: desc,
	}
}

func ParseCode(err error) Code {
	var se *SignError
	if errors.As(err, &se) {
		return codeMap[se.code]
	}
	return Code{}
}
