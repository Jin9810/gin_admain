package core

import (
	"fmt"
	"gin-vue-admin-STL/core/internal"
	"gin-vue-admin-STL/global"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	// 设置配置路径
	var config string
	if len(path) == 0 {
		config = internal.ConfigDefault
	} else {
		config = path[0]
	}

	// 读取配置
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 加载配置
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}

	return v
}
