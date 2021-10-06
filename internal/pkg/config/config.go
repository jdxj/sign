package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Bot       Bot       `yaml:"bot"`
	Logger    Logger    `yaml:"logger"`
	User      []User    `yaml:"user"`
	APIServer APIServer `yaml:"api_server"`
	Storage   Storage   `yaml:"storage"`
	DB        DB        `yaml:"db"`
	RDB       RDB       `yaml:"rdb"`
	RPC       RPC       `yaml:"rpc"`
	Rabbit    Rabbit    `yaml:"rabbit"`
	Secret    Secret    `yaml:"secret"`
}

type Bot struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}

type Logger struct {
	Path string `yaml:"path"`
	Mode string `yaml:"mode"`
}

// User 保留用
type User struct {
}

type APIServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
}

type Storage struct {
	Path string `yaml:"path"`
}

type DB struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Dbname string `yaml:"dbname"`
}

type RDB struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}

type RPC struct {
	EtcdAddr    string `yaml:"etcd_addr"`
	CrontabPort int    `yaml:"crontab_port"`
	SecretPort  int    `yaml:"secret_port"`
	UserPort    int    `yaml:"user_port"`
	NoticePort  int    `yaml:"notice_port"`
}

type Rabbit struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type Secret struct {
	Key string `yaml:"key"`
}

func ReadConfigs(path string) Root {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	r := Root{}
	err = yaml.Unmarshal(data, &r)
	if err != nil {
		log.Fatalln(err)
	}
	return r
}
