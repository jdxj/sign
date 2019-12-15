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

	"github.com/PuerkitoBio/goquery"
)

// todo: 刷 "抽 vip"
// todo: http://h5.m.iqiyi.com/integralh5/lottery/index?lotteryCode=m012801&ext=true&appKey=lottery_h5&hideshare=0&from=login&spm=31819.1.1.1&_from=share&p1=2_22_222&social_platform=link

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

type ToucherIQiYi struct {
	conf *config.IQiYiConf

	loginURL         string
	loginVerifyKey   string
	loginVerifyValue string
	signURL          string
	hotSpotBrowseURL string
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
	// 模拟浏览热点页
	hotSpotBrowseURL, err := tou.realHotSpotBrowseURL(tou.client.Jar.Cookies(tou.cookieURL))
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	req, err := utils.NewRequestWithUserAgent("GET", hotSpotBrowseURL, nil)
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

	if cip.Code != "A00000" {
		log.MyLogger.Debug("%s browse hot spot failed", log.Log_IQiYi)
		return false
	}

	// 领取浏览热点页奖励
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

	resp, err = tou.client.Do(req)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	cip, err = parseCheckInResp(data)
	if err != nil {
		log.MyLogger.Error("%s %s", log.Log_IQiYi, err)
		return false
	}

	if cip.Code == "A00000" {
		return true
	}
	return false
}

var (
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

func (tou *ToucherIQiYi) realHotSpotBrowseURL(cookies []*http.Cookie) (string, error) {
	realURL := tou.hotSpotBrowseURL + "?"

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
