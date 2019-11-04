package _8pic

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func new58Pic() (*Touch58pic, error) {
	t := &Touch58pic{
		cookies:            "qt_visitor_id=%22e035443acae82b92c2f2697d82c80ab7%22; qt_type=0; qtjssdk_2018_cross_new_user=1; qiantudata2018jssdkcross=%7B%22distinct_id%22%3A%2216e36b9c044ef-02d9d1588d8042-15291003-2073600-16e36b9c045314%22%2C%22props%22%3A%7B%22latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22latest_referrer%22%3A%22%22%2C%22latest_referrer_host%22%3A%22%22%2C%22latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D; loginTime=4; message2=1; imgCodeKey=%221501f8da399effaf4ec3205d2bc178fe%22; FIRSTVISITED=1572876174.641; loginBackUrl=%22https%3A%5C%2F%5C%2Fwww.58pic.com%5C%2F%22; popupShowNum=NaN; Hm_lvt_644763986e48f2374d9118a9ae189e14=1571093690,1571566396,1572527074,1572876175; risk_forbid_login_uid=%2257749258%22; auth_id=%2257749258%7C546L5Y%2BL5by6%7C1573480989%7Ceae3807db5907713083fe142939f5b62%22; success_target_path=%22https%3A%5C%2F%5C%2Fwww.58pic.com%5C%2F%22; sns=%7B%22token%22%3A%7B%22access_token%22%3A%2227_OZZ1mOfOU5ONPhyUCC6gnEKXrEneNKbgQXWJDQ6OeJE4lWu4sC9S1t_fWzH0gSqmUqf971_XzVDonDq6J155oCPL1k-nrEQyNYQrElc2PiQ%22%2C%22expires_in%22%3A7200%2C%22refresh_token%22%3A%2227_Rqc75NZc7ML7V2XyidW-mo7P4pdjLynWD175IAPo0-Ku0gv7SFMspBah9w7zYIx4MGsat7i_ngTg-bunE-GgnAMbxDzsud5TTvXSv-oxSDY%22%2C%22openid%22%3A%22oZw9ptwSt0KRdQjBpE4VmkiIEk6Q%22%2C%22scope%22%3A%22snsapi_login%22%2C%22unionid%22%3A%22oe6yuwy4G0SgXKAxkZxWYAf2_Ndw%22%7D%2C%22type%22%3A%22weixin%22%7D; ssid=%225dc02f9d4a8336.86641221%22; last_login_type=2; qt_risk_visitor_id=%22032df0b50c6db7e21a8e53d201409d80%22; newbieTask=%22%7B%5C%22is_login%5C%22%3A%5C%221%5C%22%2C%5C%22is_search%5C%22%3A%5C%220%5C%22%2C%5C%22is_download%5C%22%3A%5C%220%5C%22%2C%5C%22is_keep%5C%22%3A%5C%220%5C%22%2C%5C%22login_count%5C%22%3A%5C%222%5C%22%2C%5C%22upload_material%5C%22%3A%5C%220%5C%22%2C%5C%22before_login_time%5C%22%3A%5C%221570464000%5C%22%2C%5C%22is_task_complete%5C%22%3A%5C%220%5C%22%2C%5C%22task1%5C%22%3A%5C%220%5C%22%2C%5C%22task2%5C%22%3A%5C%220%5C%22%2C%5C%22task3%5C%22%3A%5C%220%5C%22%7D%22; _is_pay=0; _auth_dl_=NTc3NDkyNTh8MTU3MzQ4MDk4OXwzZjVkYjA3NjhhNjgzY2E2NGQyM2Y0ZjhkOTdlMGM0MA%3D%3D; qt_uid=%2257749258%22; censor=%2220191104%22; han_data_is_pay:57749258=%222%22; ISREQUEST=1; WEBPARAMS=is_pay=0; Hm_lpvt_644763986e48f2374d9118a9ae189e14=1572876191; qt_utime=1572876198",
		loginURL:           "https://www.58pic.com/index.php?m=IntegralMall",
		verifyKey:          ".cs-ul3-li1",
		verifyValue:        "",
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

	pic.Sign()
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
