package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"rz/global"
)

func Viper() *viper.Viper {

	path := "config/config.yaml"

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	//配置文件改动后操作
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ",e.Name)
		//将修改后的配置内容赋值到全局变量
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	//将配置内容赋值到全局变量
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
