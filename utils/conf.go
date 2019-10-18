package utils

import (
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

const (
	Conf_StudyGolang = "studygolang.com"
)

func Conf(prefix string, keys ...string) (res []string) {
	cfg := loadInI()
	for _, key := range keys {
		value := cfg.Section(prefix).Key(key).String()
		res = append(res, value)
	}

	return
}

type KeyValue struct {
	K string
	V string
}

// ConfAll 用于读取指定 section 的所有 key-value
func ConfAll(prefix string) []*KeyValue {
	cfg := loadInI()

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

func EmptyFunc() {
	time.Sleep(10*time.Second)
}

var (
	cfgMutex sync.Mutex
	iniCfg *ini.File
)

func loadInI() *ini.File {
	if iniCfg != nil {
		return iniCfg
	}

	var err error
	cfgMutex.Lock()
	defer cfgMutex.Unlock()

	iniCfg, err = ini.Load("sign.ini")
	if err != nil {
		panic(err)
	}
	return iniCfg
}