package utils

import (
	"fmt"
	"go.uber.org/zap"
	"rz/global"
	"time"
)

/**
	认证流程
 */
func DoRz() {
	fmt.Println("重新认证....")
	return
	now := time.Now().Format("2006-01-02 15:04:05")
	ip := global.CONFIG.Url.LocalIp
	info := "执行时间 : " + now + "; 本机IP : " + ip
	global.ZAP.Info("重新认证流程...", zap.Any("info", info))
	GetCsrfToken()
	DoNat()
	DoLogin()
	//认证成功, 更新配置文件
	SetCounter()
}
