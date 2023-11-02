package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

import "log"

func Test(t *testing.T) {
	// 设置配置文件的名称和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
	duration := viper.GetInt32("limit.duration")
	fmt.Println(duration)

	requests := viper.GetInt32("limit.requests")
	fmt.Println(requests)
}
