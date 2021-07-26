package configs

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Bot  Bot    `yaml:"bot"`
	User []User `yaml:"user"`
}

type Bot struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}

type User struct {
	ID     string `yaml:"id"`
	Domain string `yaml:"domain"`
	Key    string `yaml:"key"`
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