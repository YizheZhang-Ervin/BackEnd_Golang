package myconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

func Connect(configName string, configType string, configPath string) *viper.Viper {
	config := viper.New()
	if len(configPath) == 0 {
		configPath = "../configs/"
	}
	if len(configName) == 0 {
		configName = "application"
	}
	if len(configType) == 0 {
		configType = "json"
	}
	config.AddConfigPath(configPath)
	config.SetConfigName(configName)
	config.SetConfigType(configType)

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	return config
}
