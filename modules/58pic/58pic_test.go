package _8pic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	conf2 "sign/utils/conf"
	"testing"
)

func new58Pic() (*Toucher58pic, error) {
	cfg := &conf2.Pic58Conf{
		Cookies: "qt_visitor_id=%22ed68d3e81d81643fb1fe3b2aad3d1039%22; qt_type=0; message2=1; FIRSTVISITED=1575783713.044; qt_risk_visitor_id=%22fb4a220b368dc46c2b1b621b864cd843%22; ISREQUEST=1; WEBPARAMS=is_pay=0; qiantudata2018jssdkcross=%7B%22distinct_id%22%3A%2216ee407441c3d7-04014ea2d524d4-31760856-2073600-16ee407441d5d5%22%2C%22props%22%3A%7B%22latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22latest_referrer%22%3A%22%22%2C%22latest_referrer_host%22%3A%22%22%2C%22latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D; loginTime=15; imgCodeKey=%22ed9af1ec0b2b29244a0317883f598995%22; Hm_lvt_644763986e48f2374d9118a9ae189e14=1576161039,1576166356,1576244940,1576395670; loginBackUrl=%22https%3A%5C%2F%5C%2Fwww.58pic.com%5C%2F%22; popupShowNum=NaN; originUrl=https%3A%2F%2Fwww.58pic.com%2Flogin; risk_forbid_login_uid=%2257749258%22; auth_id=%2257749258%7C546L5Y%2BL5by6%7C1577000491%7C5d8a4664173cd73922b23a5509a2972c%22; success_target_path=%22https%3A%5C%2F%5C%2Fwww.58pic.com%5C%2F%22; sns=%7B%22token%22%3A%7B%22access_token%22%3A%2228_JZ-fvVo_AQzzgNnQFQj6z3CwHh7WS7TNmQnnyRbOjLW850i_CYlvurT-YM8CSn4kVHutpnTBrDfmqLfM0i9ajuRk_9p6hXeR2L81eFavHuY%22%2C%22expires_in%22%3A7200%2C%22refresh_token%22%3A%2228_JKtlEFEuPuTFiaGH46opW5qOLNmGMiat9yaTH5zWa8_MPjcr2CLbJ3IGF-TT4i6WkCA1Vz-z-T90aFBw62tFGZ_ljLfGYVsex_uH0r27c0g%22%2C%22openid%22%3A%22oZw9ptwSt0KRdQjBpE4VmkiIEk6Q%22%2C%22scope%22%3A%22snsapi_login%22%2C%22unionid%22%3A%22oe6yuwy4G0SgXKAxkZxWYAf2_Ndw%22%7D%2C%22type%22%3A%22weixin%22%7D; ssid=%225df5e3ab4deed7.57868370%22; last_login_type=2; newbieTask=%22%7B%5C%22is_login%5C%22%3A%5C%221%5C%22%2C%5C%22is_search%5C%22%3A%5C%220%5C%22%2C%5C%22is_download%5C%22%3A%5C%220%5C%22%2C%5C%22is_keep%5C%22%3A%5C%220%5C%22%2C%5C%22login_count%5C%22%3A%5C%222%5C%22%2C%5C%22upload_material%5C%22%3A%5C%220%5C%22%2C%5C%22before_login_time%5C%22%3A%5C%221570464000%5C%22%2C%5C%22is_task_complete%5C%22%3A%5C%220%5C%22%2C%5C%22task1%5C%22%3A%5C%220%5C%22%2C%5C%22task2%5C%22%3A%5C%220%5C%22%2C%5C%22task3%5C%22%3A%5C%220%5C%22%7D%22; _is_pay=0; _auth_dl_=NTc3NDkyNTh8MTU3NzAwMDQ5MXw2ZjkzOTgzMGUyM2U2ZWE1MmQ1ZWQyNzU5MGNiMjFhOQ%3D%3D; qt_uid=%2257749258%22; censor=%2220191215%22; han_data_is_pay:57749258=2; Hm_lpvt_644763986e48f2374d9118a9ae189e14=1576395695; qt_utime=1576395700",
	}

	t := &Toucher58pic{
		conf:               cfg,
		loginURL:           "https://www.58pic.com/index.php?m=IntegralMall",
		verifyKey:          ".cs-ul3-li1",
		verifyReverseValue: "我的积分:--",
		signDataURL:        "https://www.58pic.com/index.php?m=jifenNew&a=getTreeActivity",
		signURL:            "https://www.58pic.com/index.php?m=signin&a=addUserSign&time=",
		client:             &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

func TestTouch58pic_Login(t *testing.T) {
	pic, err := new58Pic()
	if err != nil {
		panic(err)
	}

	isBoot := pic.Boot()
	if !isBoot {
		fmt.Println("boot fail")
		return
	}
	isLogin := pic.Login()
	if !isLogin {
		fmt.Println("login fail")
		return
	}

	sign := pic.Sign()
	if sign {
		fmt.Println("sign success")
	} else {
		fmt.Println("sign fail")
	}
}

type People struct {
	Name string
}

func (p *People) PP() {
	fmt.Println(p.Name)
}

func TestChangePointer(t *testing.T) {
	peo := &People{Name: "123"}
	defer peo.PP()

	peo = &People{Name: "456"}
	defer peo.PP()
}

func TestJsonUnmarshal(t *testing.T) {
	str := `{"status":"1","type":1,"times":"5","clickNum":5,"week":"6","rewardThing":"\u660e\u65e5\u7b7e\u5230\u53ef\u83b7\u5f975\u79ef\u5206","valueVoucher":{"activity_name":"\u7b7e\u5230\u7d2f\u8ba1\u5956\u52b1-5\u5929","id":"94","coupon_name":"\u7b7e\u5230\u5956\u52b1-5\u5143\u6ee1\u51cf\u5238","creator":"\u8d75\u96f7","discount":"0.00","end_time":"2019-12-15","full_price":"27.00","issue_time":"0","label":"\u6ee1\u51cf\u5238","reduce_price":"5.00","relative_time":"2","start_time":"1574352026","type":"1"}}`
	sign := &sign{}

	if err := json.Unmarshal([]byte(str), &sign); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", sign)
}
