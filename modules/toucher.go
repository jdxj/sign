package modules

import (
	"net/http"
)

type Site int

const (
	Pic58 Site = iota
	StudyGolang
	Bilibili
)

// todo: 循环引用问题
//func NewToucher(cfg *modules.ToucherCfg) (modules.Toucher, error) {
//	var t modules.Toucher
//	var err error
//
//	switch cfg.Site {
//	case Pic58:
//		t, err = pic.New58Pic(cfg)
//	case StudyGolang:
//	case Bilibili:
//	}
//
//	if err != nil {
//		return nil, err
//	}
//	return t, nil
//}

type Mode int

const (
	UsePass Mode = iota + 1
	UseCookie
)

// 签到通用流程:
//     1. 构造请求 (账户名密码, cookie 等)
//     2. 访问抓取网页 (看是否登录成功)
//     3. 构造签到数据 (一些签到请求需要特殊数据, 应先获取)
//     4. 执行签到 (同时验证是否签到成功)
type Toucher interface {
	// Login 可能需要使用用户名密码或者 cookie 方式登录,
	// 其返回值 http.Cookie 不仅返回登录所使用的 cookie,
	// 还新增了 http.Response 收到的 cookie.
	// 如果 error != nil, 则没必要调用 Sign().
	Login() ([]*http.Cookie, error)
	Sign() bool
}

type ToucherCfg struct {
	Site Site // 要签到的网站
	Mode Mode

	UserName string
	Password string
	Cookies  string // 用 "key=value; key=value" 表示的字符串

	LoginURL           string // 用于验证是否登录成功所要抓取的网页
	VerifyKey          string // 指定要抓取得属性, 比如 class, li 等 html 标签或属性
	VerifyValue        string // 当要抓取的属性等于 VerifyValue 时, 判断为登录成功
	VerifyReverseValue string // 当要抓取的属性等于 VerifyValue 时, 判断为登录失败
	SignURL            string // 执行签到所要访问的链接
}

//func NewToucher(cfg *ToucherCfg) *Toucher {
//	t := &Toucher{
//		cfg: cfg,
//	}
//	return t
//}
