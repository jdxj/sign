package bilibili

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func newTouchBilibili() (*TouchBilibili, error) {
	t := &TouchBilibili{
		cookies:     "_uuid=C52E52BE-5D02-D899-B8A1-6EA59C48E68537621infoc; buvid3=760E75A7-3BE5-466C-9D36-D105DEDA550D190946infoc; LIVE_BUVID=AUTO4515713108385339; sid=89mn1hef; DedeUserID=98634211; DedeUserID__ckMd5=70cb5476ec3f0977; SESSDATA=aab9cad1%2C1573902850%2C06835ba1; bili_jct=d6c6abce34fcf6c6bcd4be93ce0dc455; bp_t_offset_98634211=318145776854474289",
		loginURL:    "https://space.bilibili.com/98634211",
		verifyKey:   "title",
		verifyValue: "王者王尼玛的个人空间 - 哔哩哔哩 ( ゜- ゜)つロ 乾杯~ Bilibili",
		client:      &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

func TestNewTouchBilibili(t *testing.T) {
	tou, _ := newTouchBilibili()

	if !tou.Boot() {
		fmt.Println("boot fail")
		return
	}

	if !tou.Login() {
		fmt.Println("login fail")
		return
	}
}
