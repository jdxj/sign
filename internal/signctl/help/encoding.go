package help

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

var (
	// AvailableFormat 用于当 Get/Create 资源时使用的格式
	AvailableFormat = map[string]struct{}{
		"":     {}, // 不处理
		"json": {},
	}
)

// Formatter 不同 kind (proto crontab.Kind) 对 key 的解释是不同的, 所以需要编码/解码成不同形式.
// 具体对 key 的解释逻辑需要参考 executor/task 的实现.
type Formatter interface {
	Encode(string) (string, error)
	Decode(string) (string, error)
}

// formatterNone 不处理
type formatterNone struct{}

func (n *formatterNone) Encode(raw string) (string, error) {
	return raw, nil
}

func (n *formatterNone) Decode(format string) (string, error) {
	return format, nil
}

// formatterJson
type formatterJson struct{}

// Encode 将 'key=value;key1=value1' 转换成 '{"key":"value","key1":"value1"}' 形式
func (j *formatterJson) Encode(raw string) (string, error) {
	req, _ := http.NewRequestWithContext(context.Background(), "", "", nil)
	req.Header.Add("Cookie", raw)
	cookies := req.Cookies()
	jar, _ := cookiejar.New(nil)
	jar.
}
