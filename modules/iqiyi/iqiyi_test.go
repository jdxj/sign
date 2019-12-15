package iqiyi

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"sign/utils"
	"sign/utils/conf"
	"testing"
)

func newIQiYi() (*ToucherIQiYi, error) {
	cfg := &conf.IQiYiConf{
		Name:        "nil",
		Cookies:     "QC005=6ce9645fc750b1b85204d0156eb4307d; QC173=0; QC006=mlhj5zc3cy6775hgw6vwbnuo; __uuid=ad1e0432-4644-d2ce-1b93-f8025144634d; QC021=%5B%7B%22key%22%3A%22%E5%A5%87%E8%91%A9%E8%AF%B4%22%7D%5D; QC124=1%7C0; QP001=1; P00004=.1573392534.97537521b9; QCUser=true; QC118=%7B%22color%22%3A%22FFFFFF%22%2C%22channelConfig%22%3A0%7D; QC159=%7B%22color%22%3A%22FFFFFF%22%2C%22channelConfig%22%3A1%2C%22speed%22%3A13%2C%22density%22%3A30%2C%22opacity%22%3A86%2C%22isFilterColorFont%22%3A1%2C%22proofShield%22%3A0%2C%22forcedFontSize%22%3A24%2C%22isFilterImage%22%3A1%2C%22isOpen%22%3A1%2C%22hadTip%22%3A1%2C%22hideRoleTip%22%3A1%7D; QP007=3720; T00404=17ca5c5c4233e4d8fa4297392ae382a9; P00001=cfjk22kzErU9qHxttB8WpNCo2nEai81WJ35zm3dy0HwzkNKMuLe2hGOGzm28qdZyfu2If6; P00003=1813046165; P00010=1813046165; P00007=cfjk22kzErU9qHxttB8WpNCo2nEai81WJ35zm3dy0HwzkNKMuLe2hGOGzm28qdZyfu2If6; P00PRU=1813046165; P00002=%7B%22uid%22%3A%221813046165%22%2C%22pru%22%3A1813046165%2C%22user_name%22%3A%2213394609376%22%2C%22nickname%22%3A%22%5Cu7528%5Cu62376c10e395%22%2C%22pnickname%22%3A%22%5Cu7528%5Cu62376c10e395%22%2C%22type%22%3A11%2C%22email%22%3A%22%22%7D; P000email=13394609376; QC160=%7B%22u%22%3A%2213394609376%22%2C%22lang%22%3A%22%22%2C%22local%22%3A%7B%22name%22%3A%22%E4%B8%AD%E5%9B%BD%E5%A4%A7%E9%99%86%22%2C%22init%22%3A%22Z%22%2C%22rcode%22%3A48%2C%22acode%22%3A%2286%22%7D%2C%22type%22%3A%22p1%22%7D; QC170=0; Hm_lvt_53b7374a63c37483e5dd97d78d9bb36e=1573392504,1575769253,1575896287; QC008=1573373587.1575769257.1575896292.4; P01010=1575907200; __dfp=a113014f0f0d28461faebfc0caccc23502e77febc919f8aa05f46cb67ffc7065cb@1577065257365@1575769258365",
		CheckInSign: "2470cb8201ddc3c8e1e5881d1f61f5d9",
		HotSpotSign: "a0212558abb20ff13a02ee5cebb36803",
		To:          "",
	}
	t := &ToucherIQiYi{
		conf:             cfg,
		loginURL:         "http://www.iqiyi.com/u/point",
		loginVerifyKey:   ".read-title-bd",
		loginVerifyValue: "做任务领VIP",
		signURL:          "http://community.iqiyi.com/openApi/score/add",
		hotSpotBrowseURL: "https://community.iqiyi.com/openApi/task/complete",
		hotSpotSignURL:   "http://community.iqiyi.com/openApi/score/getReward",
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
}

func TestRealSignURL(t *testing.T) {
	tou, _ := newIQiYi()
	cookies, err := utils.StrToCookies(tou.conf.Cookies, "mock domain")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	url, err := tou.realSignURL(cookies)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(url)
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
