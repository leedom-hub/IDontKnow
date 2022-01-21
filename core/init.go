package core

import (
	"fmt"
	"go.uber.org/zap"
	"rz/global"
	"rz/utils"
	"time"
)

func InitInfo() {
	fmtInfo := `
	==============================WARNING==============================
		        此程序只适用于测试, 请勿拿来做其他用途.
				后果自行承担.
	===================================================================
`
	fmt.Print(fmtInfo)
	now := time.Now().Format("2006-01-02 15:04:05")
	ip := global.CONFIG.Url.LocalIp
	info := "执行时间 : " + now + "; 本机IP : " + ip
	global.ZAP.Info("开始认证...", zap.Any("info", info))
	utils.GetCsrfToken()
	utils.DoNat()
	utils.DoLogin()
	//认证成功, 更新配置文件
	utils.SetCounter()
}

