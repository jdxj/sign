package modules

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
	Boot() bool
	Login() bool
	Sign() bool
}

//func NewToucher(cfg *ToucherCfg) *Toucher {
//	t := &Toucher{
//		cfg: cfg,
//	}
//	return t
//}
