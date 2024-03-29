package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Bot    Bot    `yaml:"bot"`
	Logger Logger `yaml:"logger"`
	App    App    `yaml:"app"`
	DB     DB     `yaml:"db"`
	RDB    RDB    `yaml:"rdb"`
	Rabbit Rabbit `yaml:"rabbit"`
	Secret Secret `yaml:"secret"`
	Etcd   Etcd   `yaml:"etcd"`
}

type Bot struct {
	Token string `yaml:"token"`
}

type Logger struct {
	Path string `yaml:"path"`
}

type App struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
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

type Rabbit struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type Secret struct {
	Key string `yaml:"key"`
}

type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
	Ca        string   `yaml:"ca"`
	Cert      string   `yaml:"cert"`
	Key       string   `yaml:"key"`
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
