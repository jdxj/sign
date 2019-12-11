package iqiyi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	config "sign/utils/conf"
	"sign/utils/log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// sing2: a0212558abb20ff13a02ee5cebb36803
// http://community.iqiyi.com/openApi/score/getReward
// authCookie: cfjk22kzErU9qHxttB8WpNCo2nEai81WJ35zm3dy0HwzkNKMuLe2hGOGzm28qdZyfu2If6
// userId: 1813046165
// channelCode: paopao_pcw
// agenttype: 1
// agentversion: 0
// appKey: basic_pcw
// appver: 0
// srcplatform: 1
// typeCode: point
// verticalCode: iQIYI
// scoreType: 1
// user_agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36
// dfp: a113014f0f0d28461faebfc0caccc23502e77febc919f8aa05f46cb67ffc7065cb
// sign: a0212558abb20ff13a02ee5cebb36803
// callback: cb

// 浏览热点
// Request URL: https://community.iqiyi.com/openApi/task/complete?authCookie=5errXg8Jd6m2OZHPkc3WT0kwcQcGm2Ym245f31e33qhg7iQm3nODLm2CkV7cFIm2l3tEbTb1d9&userId=1813046165&channelCode=paopao_pcw&agenttype=1&agentversion=0&appKey=basic_pcw&appver=0&srcplatform=1&typeCode=point&verticalCode=iQIYI&scoreType=1&sign=560d792c3004a0f6dce94d5205fdaf51&callback=cb
// Referer: https://www.iqiyi.com/u/point?vfrm=pcw_home&vfrmblk=A&vfrmrst=803141_points
// authCookie: 5errXg8Jd6m2OZHPkc3WT0kwcQcGm2Ym245f31e33qhg7iQm3nODLm2CkV7cFIm2l3tEbTb1d9
// userId: 1813046165
// channelCode: paopao_pcw
// agenttype: 1
// agentversion: 0
// appKey: basic_pcw
// appver: 0
// srcplatform: 1
// typeCode: point
// verticalCode: iQIYI
// scoreType: 1
// sign: 560d792c3004a0f6dce94d5205fdaf51
// callback: cb
//
// resp
// try{cb({
// "code" : "A00000",
// "message" : "成功执行.",
// "data" : {
// "dayCompleteLimit" : 1,
// "weekCompleteCount" : 3,
// "monthCompleteCount" : 5,
// "weekGetRewardCount" : 0,
// "verticalCode" : "iQIYI",
// "userId" : 1813046165,
// "totalGetRewardCount" : 0,
// "typeCode" : "point",
// "monthCompleteLimit" : 0,
// "monthGetRewardCount" : 0,
// "dayCompleteCount" : 1,
// "weekCompleteLimit" : 0,
// "dayGetRewardCount" : 0,
// "totalCompleteCount" : 5,
// "cooldown" : 0,
// "totalCompleteLimit" : 0,
// "channelCode" : "paopao_pcw"
// }
// })}catch(e){}
func NewIQiYiFromApi(conf *config.IQiYiConf) (*ToucherIQiYi, error) {
	if conf == nil {
		return nil, fmt.Errorf("invalid cfg")
	}

	t := &ToucherIQiYi{
		conf:             conf,
		loginURL:         "http://www.iqiyi.com/u/point",
		loginVerifyKey:   ".read-title-bd",
		loginVerifyValue: "做任务领VIP",
		signURL:          "http://community.iqiyi.com/openApi/score/add",
		hotSpotURL:       "http://www.iqiyi.com/feed/",
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

type ToucherIQiYi struct {
	conf *config.IQiYiConf

	loginURL         string
	loginVerifyKey   string
	loginVerifyValue string
	signURL          string
	hotSpotURL       string
	hotSpotSignURL   string

	client    *http.Client
	cookieURL *url.URL
}

func (tou *ToucherIQiYi) Name() string {
	return tou.conf.Name
}
func (tou *ToucherIQiYi) Email() string {
	return tou.conf.To
}

func (tou *ToucherIQiYi) Boot() bool {
	cookies, err := utils.StrToCookies(tou.conf.Cookies, utils.IQiYiCookieDomain)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	cookieURL, err := url.Parse(utils.IQiYiCookieURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}
	tou.cookieURL = cookieURL

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

func (tou *ToucherIQiYi) Login() bool {
	resp, err := tou.client.Get(tou.loginURL)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}
	rawText := doc.Find(tou.loginVerifyKey).Text()
	if strings.TrimSpace(rawText) == tou.loginVerifyValue {
		return true
	}
	return false
}

func (tou *ToucherIQiYi) Sign() bool {
	// 保证都签一次
	ciStat := tou.checkIn()
	hsStat := tou.hotSpot()
	if ciStat || hsStat {
		return true
	}

	return false
}

func (tou *ToucherIQiYi) checkIn() bool {
	signURL, err := tou.realSignURL(tou.client.Jar.Cookies(tou.cookieURL))
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	req, err := utils.NewRequestWithUserAgent("GET", signURL, nil)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	resp, err := tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	cip, err := parseCheckInResp(data)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	if cip.Code == "A00000" {
		return true
	}
	return false
}

func (tou *ToucherIQiYi) hotSpot() bool {
	// todo: 浏览失败, 需要修复
	req, err := utils.NewRequestWithUserAgent("GET", tou.hotSpotURL, nil)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	// 模拟浏览热点页
	_, _ = tou.client.Do(req)
	time.Sleep(time.Second)

	hotSpotSignURL, err := tou.realhotSpotSignURL(tou.client.Jar.Cookies(tou.cookieURL))
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	req, err = utils.NewRequestWithUserAgent("GET", hotSpotSignURL, nil)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	resp, err := tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	cip, err := parseCheckInResp(data)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	if cip.Code == "A00000" {
		return true
	}
	return false
}

// authCookie: cfjk22kzErU9qHxttB8WpNCo2nEai81WJ35zm3dy0HwzkNKMuLe2hGOGzm28qdZyfu2If6
// userId: 1813046165
// channelCode: sign_pcw
// agenttype: 1
// agentversion: 0
// appKey: basic_pcw
// appver: 0
// srcplatform: 1
// typeCode: point
// verticalCode: iQIYI
// scoreType: 1
// user_agent: Mozilla%2F5.0%20(X11%3B%20Linux%20x86_64)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Chrome%2F78.0.3904.108%20Safari%2F537.36
// dfp: a113014f0f0d28461faebfc0caccc23502e77febc919f8aa05f46cb67ffc7065cb
// sign: 2470cb8201ddc3c8e1e5881d1f61f5d9
// callback: cb
//
// queryParams 是签到时 (signURL) 所使用的 url 查询参数.
// 其中注释掉的使用固定值.
var (
	// "channelCode" 根据签到类型不同, 有不同的值
	//     sign    : sign_pcw
	//     hot spot: paopao_pcw

	// todo: "sign" 需要得知如何计算
	queryParamsCookie = map[string]string{
		"P00001": "authCookie",
		"P00003": "userId",
		"__dfp":  "dfp",
	}

	queryParamsFixed = map[string]string{
		"agenttype":    "1",
		"agentversion": "0",
		"appKey":       "basic_pcw",
		"appver":       "0",
		"srcplatform":  "1",
		"typeCode":     "point",
		"verticalCode": "iQIYI",
		"scoreType":    "1",
		"user_agent":   "Mozilla%2F5.0%20(X11%3B%20Linux%20x86_64)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Chrome%2F78.0.3904.108%20Safari%2F537.36",
		"callback":     "cb",
	}
)

func (tou *ToucherIQiYi) realSignURL(cookies []*http.Cookie) (string, error) {
	realSignURL := tou.signURL + "?"
	cm := utils.CookieArrayToMap(cookies)

	for qpRawK, qpK := range queryParamsCookie {
		if cookie, ok := cm[qpRawK]; ok {
			if qpRawK == "__dfp" {
				realSignURL += qpK + "=" + getDFP(cookie.Value) + "&"
			} else {
				realSignURL += qpK + "=" + cookie.Value + "&"
			}
		} else {
			return "", fmt.Errorf("not found query param: %s", qpK)
		}
	}

	for qpK, qpV := range queryParamsFixed {
		realSignURL += qpK + "=" + qpV + "&"
	}
	realSignURL += "channelCode=sign_pcw&"
	realSignURL += "sign=" + tou.conf.CheckInSign
	return realSignURL, nil
}

func (tou *ToucherIQiYi) realhotSpotSignURL(cookies []*http.Cookie) (string, error) {
	realURL := tou.hotSpotSignURL + "?"

	cm := utils.CookieArrayToMap(cookies)

	for qpRawK, qpK := range queryParamsCookie {
		if cookie, ok := cm[qpRawK]; ok {
			if qpRawK == "__dfp" {
				realURL += qpK + "=" + getDFP(cookie.Value) + "&"
			} else {
				realURL += qpK + "=" + cookie.Value + "&"
			}
		} else {
			return "", fmt.Errorf("not found query param: %s", qpK)
		}
	}

	for qpK, qpV := range queryParamsFixed {
		realURL += qpK + "=" + qpV + "&"
	}
	realURL += "channelCode=paopao_pcw&"
	realURL += "sign=" + tou.conf.HotSpotSign
	return realURL, nil
}

func getDFP(v string) string {
	vs := strings.Split(v, "@")
	if len(vs) < 1 {
		return ""
	}
	return vs[0]
}

// try{cb({
// "code" : "A00000",
// "message" : "成功执行.",
// "data" : [ {
// "code" : "A0002",
// "trdetailList" : null,
// "curTRDetail" : null,
// "trlotDetailList" : null,
// "nextTRLotDetail" : null,
// "signDayForCycle" : 0,
// "message" : "任务次数已经到达上限",
// "curTRLotDetail" : null,
// "nextTRDetail" : null,
// "typeCode" : "point",
// "continuousScore" : 0,
// "score" : 0,
// "continuousValue" : 0,
// "rewardCode" : null
// } ]
// })}catch(e){}
type checkInResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func parseCheckInResp(resp []byte) (*checkInResp, error) {
	resp = bytes.TrimPrefix(resp, []byte("try{cb("))
	resp = bytes.TrimSuffix(resp, []byte(")}catch(e){}"))

	cip := &checkInResp{}
	if err := json.Unmarshal(resp, cip); err != nil {
		return nil, err
	}
	return cip, nil
}

// , x = {
// uid: r.getUid(),
// authCookie: r.getAuthCookies(),
// appKeyQiyi: "basic_pcw",
// secret_key_qiyi: "UKobMjDMsDoScuWOfp6F",
// needExpire: 1,
// agenttype: "1",
// agentversion: "0",
// appver: "0",
// srcplatform: "1",
// typeCode: "point",
// verticalCode: "iQIYI"
