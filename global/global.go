package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"rz/config"
)

var (
	CONFIG config.Server
	VP     *viper.Viper
	ZAP     *zap.Logger
)