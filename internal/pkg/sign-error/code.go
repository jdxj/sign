package sign_error

const (
	ErrUnknown = iota
	ErrBindReqFailed
	ErrAuthFailed
	ErrHandle
	ErrRPCRequest
	ErrInternal
	ErrInvalidParam
)

// todo: 代码自动生成
var (
	ErrMap = map[int]string{
		ErrUnknown:       "未知错误",
		ErrBindReqFailed: "解析请求错误",
		ErrAuthFailed:    "认证失败",
		ErrRPCRequest:    "服务调用失败",
		ErrInternal:      "内部错误",
		ErrInvalidParam:  "无效数据",
		ErrHandle:        "处理请求错误",
	}
)
