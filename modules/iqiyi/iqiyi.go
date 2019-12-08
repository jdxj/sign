package iqiyi

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sign/utils"
	config "sign/utils/conf"
	"sign/utils/log"
	"strings"

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

	client *http.Client
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

	tou.client.Jar.SetCookies(cookieURL, cookies)
	return true
}

// todo: 实现
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

// todo: 实现
func (tou *ToucherIQiYi) Sign() bool {
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
			realSignURL += qpK + "=" + cookie.Value + "&"
		} else {
			return "", fmt.Errorf("not found query param: %s", qpK)
		}
	}

	for qpK, qpV := range queryParamsFixed {
		realSignURL += qpK + "=" + qpV + "&"
	}
	realSignURL += "channelCode=sign_pcw&"
	realSignURL += genSignQueryParam()
	return realSignURL, nil
}

// todo: 构造 "sign" 查询参数
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
func genSignQueryParam() string {
	// 模拟的
	return "sign=a0212558abb20ff13a02ee5cebb36803"
}
