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
	APIServer APIServer `json:"api_server"`
}

type Bot struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}

type Logger struct {
	Path string `yaml:"path"`
	Mode string `yaml:"mode"`
}

type User struct {
	ID     string `yaml:"id" json:"id"`
	Domain int    `yaml:"domain" json:"domain"`
	Type   []int  `yaml:"type" json:"type"`
	Key    string `yaml:"key" json:"key"`
}

type APIServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
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
