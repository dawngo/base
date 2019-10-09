package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig *viper.Viper
}

var Conf *Config

func init() {
	Conf = New()
	Conf.LoadGlobalConfig()
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadGlobalConfig() {
	c.AppConfig = viper.New()

	//设置配置文件的名字
	c.AppConfig.SetConfigName("global")
	//添加配置文件所在的路径
	c.AppConfig.AddConfigPath("./config")
	//设置配置文件类型
	c.AppConfig.SetConfigType("yaml")
	err := c.AppConfig.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
