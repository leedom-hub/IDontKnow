package utils

import (
	"github.com/imroc/req"
	"go.uber.org/zap"
	"io/ioutil"
	"regexp"
	"rz/global"
	"strconv"
)

func GetLeftTime() int {
	get, err := req.Get(global.CONFIG.Url.LoginOutUrl)
	if err != nil {
		global.ZAP.Error("GetLeftTime err ",zap.Any("err",err))
	}
	resp := get.Response()
	defer func() {
		err2 :=resp.Body.Close()
		if err2 != nil {
			global.ZAP.Error("GetLeftTime defer err ",zap.Any("err",err))
		}
	}()

	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		global.ZAP.Info("GetLeftTime Body not close success", zap.Any("info", err3))
	}
	src := string(body)
	//先取到各个正则规则
	reStrTimer, _ := regexp.Compile(`(?m)var timer = (.\d*)`)
	reStrRegistTime, _ := regexp.Compile(`(?m)var regist_time = (.\d*)`)
	reStrNow, _ := regexp.Compile(`(?m)var now = (.\d*)`)
	//再通过规则获取字符串
	strTimer := reStrTimer.FindString(src)
	strRegistTime := reStrRegistTime.FindString(src)
	strNow := reStrNow.FindString(src)

	//再取字符串中的时间
	reTimer := regexp.MustCompile("var timer = (.\\d*)")
	reRegistTime := regexp.MustCompile("var regist_time = (.\\d*)")
	reNow := regexp.MustCompile("var now = (.\\d*)")

	timerS := reTimer.FindStringSubmatch(strTimer)
	registTimeS := reRegistTime.FindStringSubmatch(strRegistTime)
	nowS := reNow.FindStringSubmatch(strNow)

	timerSs := timerS[len(timerS)-1]
	registTimeSs := registTimeS[len(registTimeS)-1]
	nowSs := nowS[len(nowS)-1]

 	timer, _ := strconv.Atoi(timerSs)
	registTime, _ := strconv.Atoi(registTimeSs)
	now, _ := strconv.Atoi(nowSs)

	return (registTime + timer) - now
}

func SetCounter() {
	left := GetLeftTime()
	global.VP.Set("counter.left",left)
	_ = global.VP.WriteConfig()
	global.ZAP.Info("获取剩余时间, 并更新配置文件;", zap.Any("counter.left = ", left))
}


