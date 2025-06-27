package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var Cfg Config

func InitConfig() {
	viper.SetConfigName("config") // 不带扩展名
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")     // 当前目录
	viper.AddConfigPath("../..") // 适用于在 internal/db 测试运行时
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("读取配置失败: %v", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Printf("配置文件解析失败: %v", err)
	}
}
