package _8pic

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func new58Pic() (*Toucher58pic, error) {
	t := &Toucher58pic{
		cookies:            "qt_visitor_id=%2212d0ffced2e9fa6ed800126dc7b12408%22&qt_type=0&qiantudata2018jssdkcross=%7B%22distinct_id%22%3A%2216dd8d58d6715-04c694b80d406f-38677b07-2073600-16dd8d58d6859d%22%2C%22props%22%3A%7B%22latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22latest_referrer%22%3A%22%22%2C%22latest_referrer_host%22%3A%22%22%2C%22latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D&message2=1&FIRSTVISITED=1571300938.291&qt_risk_visitor_id=%220273633cb51e6a1e481121cf5785d2d4%22&ISREQUEST=1&WEBPARAMS=is_pay=0&Hm_lvt_644763986e48f2374d9118a9ae189e14=1571314095&first-charge-57749258_10=patchOne%3A2%2CtimeOne%3A17%2CtimeEnd%3A10&58pic_btn=l_h_1&last_login_type=2&qt-couponFlag-57749258=time%3A28&qt-couponFlag-undefined=time%3A31&qt-coupon-tips=0&success_target_path=%22https%3A%5C%2F%5C%2Fwww.58pic.com%5C%2F%22&loginTime=4&imgCodeKey=%2281a777f57ace12196647e60984ce31dc%22&popupShowNum=1&risk_forbid_login_uid=%2257749258%22&auth_id=%2257749258%7C546L5Y%2BL5by6%7C1573465772%7C2f4b916940fe1dafe3b5f4dabb54b1bc%22&sns=%7B%22token%22%3A%7B%22access_token%22%3A%2227_EM1gMCRh9cU24pCZfZQOgnAnzeJ4ExHq3b9S1-E1mKunA3w9E_yQA5UqHlQNeIq8AFK2rqaf_JHjZyJua43nm20KckIKNao_PRTz8YhbKIc%22%2C%22expires_in%22%3A7200%2C%22refresh_token%22%3A%2227_7MUcriNdvus4DmsMDwW_R-XEq_X_atEaABqJhjq05eL9VhN1pJ2K_04YfaVxvztPHmaOb8StxFCICzUPO_4BVQ6wHT40-dMWNlSrWGFeKvI%22%2C%22openid%22%3A%22oZw9ptwSt0KRdQjBpE4VmkiIEk6Q%22%2C%22scope%22%3A%22snsapi_login%22%2C%22unionid%22%3A%22oe6yuwy4G0SgXKAxkZxWYAf2_Ndw%22%7D%2C%22type%22%3A%22weixin%22%7D&ssid=%225dbff42cdae6e7.71941317%22&newbieTask=%22%7B%5C%22is_login%5C%22%3A%5C%221%5C%22%2C%5C%22is_search%5C%22%3A%5C%220%5C%22%2C%5C%22is_download%5C%22%3A%5C%220%5C%22%2C%5C%22is_keep%5C%22%3A%5C%220%5C%22%2C%5C%22login_count%5C%22%3A%5C%222%5C%22%2C%5C%22upload_material%5C%22%3A%5C%220%5C%22%2C%5C%22before_login_time%5C%22%3A%5C%221570464000%5C%22%2C%5C%22is_task_complete%5C%22%3A%5C%220%5C%22%2C%5C%22task1%5C%22%3A%5C%220%5C%22%2C%5C%22task2%5C%22%3A%5C%220%5C%22%2C%5C%22task3%5C%22%3A%5C%220%5C%22%7D%22&_is_pay=0&_auth_dl_=NTc3NDkyNTh8MTU3MzQ2NTc3M3wyYmMwM2Q4YzNlZDU2ZWE0MGVhNjVlYjY5MTU4ZDM2Mg%3D%3D&qt_uid=%2257749258%22&censor=%2220191105%22&han_data_is_pay:57749258=2&qt_utime=1572922989&Hm_lpvt_644763986e48f2374d9118a9ae189e14=1572922991",
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
