package help

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// todo: map -> struct
var (
	// AvailableFormat 用于当 Get/Create 资源时使用的格式
	AvailableFormat = map[string]Formatter{
		"":     &formatterNone{},
		"json": &formatterJson{},
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
	res := make(map[string]string)
	header := map[string][]string{
		"Cookie": {raw},
	}
	req := &http.Request{Header: header}
	for _, cookie := range req.Cookies() {
		res[cookie.Name] = cookie.Value
	}
	d, err := json.Marshal(res)
	return string(d), err
}

// Decode 将 '{"key":"value","key1":"value1"}' 转换成 'key=value;key1=value1' 形式
func (j *formatterJson) Decode(format string) (string, error) {
	var res map[string]string
	err := json.Unmarshal([]byte(format), &res)
	if err != nil {
		return "", err
	}
	kv := make([]string, 0, len(res))
	for k, v := range res {
		kv = append(kv, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(kv, ";"), nil
}
