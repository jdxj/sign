package _8pic

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func new58Pic() (*Touch58pic, error) {
	t := &Touch58pic{
		cookies:            "",
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
