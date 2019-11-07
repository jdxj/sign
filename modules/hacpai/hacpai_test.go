package hacpai

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newToucherHacPai() (*ToucherHacPai, error) {
	tou := &ToucherHacPai{
		//username: "",
		//password: "",
		cookies:     "Hm_lvt_f241a238dc8343347478081db6c7cf5c=1573137619&_ga=GA1.2.1769622619.1573137620&_gid=GA1.2.1667096449.1573137620&_ga=GA1.2.1769622619.1573137620&_gid=GA1.2.1667096449.1573137620&Hm_lvt_f241a238dc8343347478081db6c7cf5c=1573137619&Hm_lpvt_f241a238dc8343347478081db6c7cf5c=1573137619&_gat_gtag_UA_144249821_1=1&Hm_lpvt_f241a238dc8343347478081db6c7cf5c=1573141552&symphony=b1d0d1ad98acdd5d7d846d4359048a923f932962125e7ad0c9db814d5942879467b648f693d6a11336bcaa8b1bee42097e27e816d3225de0ee60efaaeeb73c9d21a76fc165df292b14125255a4fb239d62d04bb21c4ac714e5e46b0e2af65dcdcdb31cbc50accdfe7e7d6b5c224e1b882a17c84872510af4c74abe53db0691f0530594ad52b98a8e08365e5e5f6f9e51b6a8b29db45e6046933b171df307449e288308c5e1981a2cd63b1f13d4cb75fdfc005c99d79707b1bb83b06af5a5935e9c26c45810989439346f7172293111bb8e3cf1f2ddbfff2d88d531529d21d6bc",
		loginURL:    "https://hacpai.com/member/AaronWang",
		signURL:     "https://hacpai.com/activity/checkin",
		verifyKey:   ".fn-inline",
		verifyValue: "jdxj",
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tou.client.Jar = jar
	return tou, nil
}

func TestToucherHacPai(t *testing.T) {
	tou, err := newToucherHacPai()
	if err != nil {
		panic(err)
	}

	if !tou.Boot() {
		fmt.Println("boot fail")
	}
	if !tou.Login() {
		fmt.Println("login fail")
	}
	if !tou.Sign() {

	}
}
