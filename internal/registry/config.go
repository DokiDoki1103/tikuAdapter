package registry

import (
	"bytes"
	"fmt"
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
limit:
  enable: false # 是否开启
  duration: 3  # 时间窗口为3秒
  requests: 1  # 允许用户在3秒内通过1个请求
`)

// Config get config instance
func Config() configs.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		logger.SysLog(fmt.Sprintf("无法读取配置文件: %v,将以默认配置启动", err))
		_ = viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	}

	var config configs.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		logger.SysError(fmt.Sprintf("配置文件语法错误: %v", err))
	}
	return config
}
