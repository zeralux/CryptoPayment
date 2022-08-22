package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

var appConfig *AppConfig

// copy value
func GetAppConfig() AppConfig {
	return *appConfig
}

func LoadAppConfig(configLocation string) {
	if appConfig == nil {
		// 預設
		setDefault(appConfig)
		// 添加驱动程序以支持yaml内容解析（除了JSON是默认支持，其他的则是按需使用）
		config.AddDriver(yaml.Driver)
		// 加载配置，可以同时传入多个文件
		err := config.LoadFiles(configLocation)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("config data: \n %#v\n", config.Data())
		config.BindStruct("", &appConfig)
	}
}

func setDefault(appConfig *AppConfig) {
	appConfig.service = Service{
		port: 8080,
	}
}

type AppConfig struct {
	service Service `mapstructure:"service"`
}

func (a AppConfig) Service() Service {
	return a.service
}
