package main

// 该文件用于快速的添加任务

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sign/utils/conf"
)

type Conf struct {
	StudyGolangs []conf.SGConf     `json:"studyGolangs"`
	Bilibilis    []conf.BiliConf   `json:"bilibilis"`
	Pic58s       []conf.Pic58Conf  `json:"pic58s"`
	HacPais      []conf.HacPaiConf `json:"hacPais"`
	V2exs        []conf.V2exConf   `json:"v2exs"`
	IQiYis       []conf.IQiYiConf  `json:"iQiYis"`

	Host     string `json:"host"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Resp struct {
	Msg string `json:"msg"`
}

var apis = [6]string{
	"/api/studygolang",
	"/api/bilibili",
	"/api/58pic",
	"/api/hacpai",
	"/api/v2ex",
	"/api/iqiyi",
}

func main() {
	file, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	conf := &Conf{}
	if err := json.Unmarshal(data, conf); err != nil {
		panic(err)
	}

	join := conf.UserName + ":" + conf.Password
	client := &http.Client{}
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(join))
	fmt.Println(basic)

	// studygolang
	for _, sg := range conf.StudyGolangs {
		url := conf.Host + apis[0]

		fmt.Printf("%s\n", sg.Name)
		data, err := json.Marshal(sg)
		if err != nil {
			panic(err)
		}

		send(data, url, basic, client)
	}

	for _, bl := range conf.Bilibilis {
		url := conf.Host + apis[1]

		fmt.Printf("%s\n", bl.Name)
		data, err := json.Marshal(bl)
		if err != nil {
			panic(err)
		}

		send(data, url, basic, client)
	}

	for _, pic58 := range conf.Pic58s {
		url := conf.Host + apis[2]

		fmt.Printf("%s\n", pic58.Name)
		data, err := json.Marshal(pic58)
		if err != nil {
			panic(err)
		}

		send(data, url, basic, client)
	}
	for _, hacPai := range conf.HacPais {
		url := conf.Host + apis[3]

		fmt.Printf("%s\n", hacPai.Name)
		data, err := json.Marshal(hacPai)
		if err != nil {
			panic(err)
		}

		send(data, url, basic, client)
	}
	for _, v2ex := range conf.V2exs {
		url := conf.Host + apis[4]

		fmt.Printf("%s\n", v2ex.Name)
		data, err := json.Marshal(v2ex)
		if err != nil {
			panic(err)
		}

		send(data, url, basic, client)
	}
	for _, iQiYi := range conf.IQiYis {
		url := conf.Host + apis[5]

		fmt.Printf("%s\n", iQiYi.Name)
		data, err := json.Marshal(iQiYi)
		if err != nil {
			panic(err)
		}

		send(data, url, basic, client)
	}
}

func send(data []byte, url, basic string, client *http.Client) {
	rd := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, rd)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", basic)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	stat := &Resp{}
	if err := json.Unmarshal(data, stat); err != nil {
		fmt.Printf("%s\n", data)
		panic(err)
	}

	if stat.Msg != "add success" {
		fmt.Printf("add failed: %s\n", data)
	}
}
