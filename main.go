package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"rz/core"
	"rz/global"
	"rz/utils"
	"time"
)

func main() {
	global.VP = core.Viper()
	global.ZAP = core.Zap()
	// 设置代理
	if global.CONFIG.Env.Proxy {
		_ = os.Setenv("HTTP_PROXY", "http://127.0.0.1:8888")
	}
	//初始化 进行第一次认证
	core.InitInfo()

	tick := time.Tick(1 * time.Second)
	left := global.CONFIG.Counter.Left
	global.ZAP.Info("初始化倒计时...", zap.Any("info", left))
	for  {
		left -= 1
		if left < 1 {
			//重新认证
			utils.DoRz()
			left = global.CONFIG.Counter.Left
		}else{
			//倒计时
			if global.CONFIG.Env.OutPutCount {
				fmt.Println("倒计时 : ", left)
			}
		}
		<-tick
	}

}


