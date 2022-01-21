package utils

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"go.uber.org/zap"
	"io/ioutil"
	"regexp"
	"rz/config"
	"rz/global"
	"strings"
)

var (
	CsrfToken string
)

func GetCsrfToken() {
	get, err := req.Get(global.CONFIG.Url.LoginUrl)
	if err != nil {
		global.ZAP.Error("CsrfToken err ",zap.Any("err",err))
	}
	 resp := get.Response()
	defer func() {
		err2 :=resp.Body.Close()
		if err2 != nil {
			//global.ZAP.Error("CsrfToken defer err ",zap.Any("err",err))
		}
	}()

	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		global.ZAP.Info("CsrfToken Body close fail", zap.Any("info", err3))
	}
	src := string(body)
	re, _ := regexp.Compile(`(?m){"csrf_token": (.*)`)
	tokenJson := &config.Token{}
	_ = json.Unmarshal([]byte(re.FindString(src)), tokenJson)

	CsrfToken = strings.Replace(tokenJson.CsrfToken, " ", "", -1)
	CsrfToken = strings.Replace(CsrfToken, "\n", "", -1)
	CsrfToken = strings.Trim(CsrfToken, " ")

	global.ZAP.Info("获取CsrfToken", zap.Any("info", CsrfToken))
	return
}

func DoNat() {
	myHeader := req.HeaderFromStruct(GenerateHeader())
	r, err := req.Post(global.CONFIG.Url.NatApi, myHeader, req.Param{
		"0.0.0.0":false,
		global.CONFIG.Url.LocalIp:true,
	})
	if err != nil {
		global.ZAP.Error("do Nat err ",zap.Any("err",err))
	}
	resp := r.Response()
	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			global.ZAP.Info("Nat Body close fail", zap.Any("info", err2))
		}
	}()
	_body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		global.ZAP.Info("Nat Body close fail", zap.Any("info", err3))
	}
	global.ZAP.Info("Nat接口成功!", zap.Any("info", string(_body)))
}

func GenerateHeader() config.Header {
	return config.Header{
		Accept:         global.CONFIG.Header.Accept,
		AcceptEncoding: global.CONFIG.Header.AcceptEncoding,
		AcceptLanguage: global.CONFIG.Header.AcceptLanguage,
		Connection:     global.CONFIG.Header.Connection,
		Host:           global.CONFIG.Header.Host,
		Origin:         global.CONFIG.Header.Origin,
		Referer:        global.CONFIG.Header.Referer,
		UserAgent:      global.CONFIG.Header.UserAgent,
		XCSRFToken:     CsrfToken,
		XRequestedWith: "XMLHttpRequest",
	}


}

func DoLogin() {
	if global.CONFIG.Env.Debug {
		req.Debug = true
	}

	myHeader := req.HeaderFromStruct(GenerateHeader())
	fmt.Println(myHeader)
	r, err := req.Post(global.CONFIG.Url.LoginApi, myHeader, req.Param{
		"user_name" : global.CONFIG.User.UserName,
		"pwd" : global.CONFIG.User.Pwd,
		"request_url" :global.CONFIG.LoginForm.RequestUrl,
		"os_name" : global.CONFIG.LoginForm.OsName,
		"browser_name" : global.CONFIG.LoginForm.BrowserName,
		"force_change" : global.CONFIG.LoginForm.ForceChange,
	})
	if err != nil {
		global.ZAP.Error("do login err ",zap.Any("err",err))
	}

	resp := r.Response()

	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			global.ZAP.Info("Body not close success", zap.Any("info", err2))
		}
	}()
	_body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		global.ZAP.Info("ReadAll Body not close success", zap.Any("info", err3))
	}

	if resp.StatusCode != 200 {
		global.ZAP.Warn("登录接口失败!", zap.Any("info", string(_body)))
		return
	}

	global.ZAP.Info("登录接口成功!", zap.Any("info", string(_body)))
}
