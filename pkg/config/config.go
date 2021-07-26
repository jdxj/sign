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
	User   []User `yaml:"user"`
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
	ID   string `yaml:"id"`
	Type int    `yaml:"type"`
	Key  string `yaml:"key"`
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
