package utils

import (
	"gopkg.in/ini.v1"
)

const (
	Conf_StudyGolang = "studygolang.com"
)

func Conf(prefix string, keys ...string) (res []string) {
	cfg, err := ini.Load("sign.ini")
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		value := cfg.Section(prefix).Key(key).String()
		res = append(res, value)
	}

	return
}

// todo: 声明全局 cfg, 不要多次读取同一配置文件
// 优点是可以更改配置再写回文件中

type KeyValue struct {
	K string
	V string
}

// ConfAll 用于读取指定 section 的所有 key-value
func ConfAll(prefix string) []*KeyValue {
	cfg, err := ini.Load("sign.ini")
	if err != nil {
		panic(err)
	}

	var kvs []*KeyValue
	sec := cfg.Section(prefix)
	for _, key := range sec.Keys() {
		kv := &KeyValue{
			K: key.Name(),
			V: key.String(),
		}

		kvs = append(kvs, kv)
	}
	return kvs
}
