package config

import (
	"github.com/spf13/viper"
)

//应用程序配置
type appConfig struct {
	//应用名称
	Name string
	//运行模式 debug， release， test
	RunMode string
	// 运行 addr
	Addr string
	// 完整 url
	URL string
	// secret key
	Key string
}

func newAppConfig() *appConfig {
	//default setting
	viper.SetDefault("APP.NAME", "gin.weibo")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP", ":8080")
	viper.SetDefault("APP.URL", "")
	viper.SetDefault("APP.KEY", "base64:O+VQ74YEigLPDzLKnh2HW/yjCdU2ON9v7xuKBgSOEAo=")

	return &appConfig{
		Name:    viper.GetString("APP.NAME"),
		RunMode: viper.GetString("APP.RUNMODE"),
		Addr:    viper.GetString("APP.ADDR"),
		URL:     viper.GetString("APP.URL"),
		Key:     viper.GetString("APP.KEY"),
	}
}
