package iqiyi

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"sign/utils/conf"
	"testing"
)

//http://community.iqiyi.com/openApi/lottery/multiDraw?userId=1813046165&authCookie=e7Fw6m1lyRAMrh0m2g72kWJm2vJ4187IeI4IFrIykzxlhO5yrfedYDtlJGfxOm1rJqZL9f26&agenttype=21&agentversion=9.3.0&srcplatform=21&appver=9.3.0&lotteryCode=m012801&times=5&appKey=lottery_h5&sign=bf49e9387a9dc627595982aa17e0ebf9
func newIQiYi() (*ToucherIQiYi, error) {
	cfg := &conf.IQiYiConf{
		Name:          "nil",
		Cookies:       ``,
		CheckInSign:   "",
		HotSpotBrowse: "",
		HotSpotSign:   "",
		To:            "",
	}
	t := &ToucherIQiYi{
		conf:             cfg,
		loginURL:         "http://www.iqiyi.com/u/point",
		loginVerifyKey:   ".read-title-bd",
		loginVerifyValue: "做任务领VIP",
		client:           &http.Client{},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	t.client.Jar = jar
	return t, nil
}

func TestToucherIQiYi(t *testing.T) {
	tou, _ := newIQiYi()

	if !tou.Boot() {
		t.Fatalf("boot failed")
	}
	if !tou.Login() {
		t.Fatalf("login failed")
	}
	if !tou.Sign() {
		t.Fatalf("check in failed")
	}
	//if !tou.hotSpot() {
	//	t.Fatalf("hot spot fail")
	//}
	//if !tou.checkIn() {
	//	t.Fatalf("check in fail")
	//}
}

func TestRealSignURL(t *testing.T) {
}

func TestParseCheckInResp(t *testing.T) {
	data := `try{cb({
  "code" : "A00000",
  "message" : "成功执行.",
  "data" : [ {
    "code" : "A0002",
    "trdetailList" : null,
    "curTRDetail" : null,
    "trlotDetailList" : null,
    "nextTRLotDetail" : null,
    "signDayForCycle" : 0,
    "message" : "任务次数已经到达上限",
    "curTRLotDetail" : null,
    "nextTRDetail" : null,
    "typeCode" : "point",
    "continuousScore" : 0,
    "score" : 0,
    "continuousValue" : 0,
    "rewardCode" : null
  } ]
})}catch(e){}`
	cip, err := parseCheckInResp([]byte(data))
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%+v\n", cip)
}

func TestRealHotSpotSignURL(t *testing.T) {
	tou, _ := newIQiYi()
	if !tou.Boot() {
		t.Fatalf("boot failed")
	}
	if !tou.Login() {
		t.Fatalf("login failed")
	}
	if !tou.hotSpot() {
		t.Fatalf("%s\n", "hot spot sign failed")
	}
}
