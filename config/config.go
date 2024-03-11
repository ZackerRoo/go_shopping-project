package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfgReader *configReader

type (
	Configuration struct {
		DatabaseSettings
		JwtSettings
	}
	// 数据库配置
	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}
	// jwt配置
	JwtSettings struct {
		SecretKey string
	}
	// reader
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

// 获得所有配置
func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	newConfigReader(configFile)
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取失败 : %s", err)
		return nil, err
	}

	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件到结构体失败 : %s", err)
		return nil, err
	}

	return configuration, err
}

// 实例化configReader
func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}
