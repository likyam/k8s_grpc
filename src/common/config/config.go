package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config Service config
type Config struct {
	Grpc     Grpc     `json:"grpc" yaml:"grpc"`
	Http     Http     `json:"http" yaml:"http"`
	Database Database `json:"database" yaml:"database"`
	Redis    Redis    `json:"redis" yaml:"redis"`
	Client   Client   `json:"client" yaml:"client"`
	Trace    Trace    `json:"trace" yaml:"trace"`
}

// NewConfig Initial service's config
func NewConfig(cfg string) *Config {

	//if cfg == "" {
	//	panic("load config file failed.config file can not be empty.")
	//}
	//
	//viper.SetConfigFile(cfg)
	//
	//// Read config file
	//if err := viper.ReadInConfig(); err != nil {
	//	panic("read config failed.[ERROR]=>" + err.Error())
	//}
	//
	//// 在读取第一个配置文件后，你可以继续读取其他配置文件
	//viper.SetConfigFile("config2.yaml")
	//if err := viper.MergeInConfig(); err != nil {
	//	fmt.Printf("无法读取配置文件: %v\n", err)
	//	return nil
	//}
	// 创建一个新的 Viper 实例
	viper := viper.New()

	// 设置要读取的配置文件的文件名（可以是相对路径或绝对路径）
	viper.SetConfigFile(cfg)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %v\n", err)
		return nil
	}

	// 在读取第一个配置文件后，你可以继续读取其他配置文件
	viper.SetConfigFile("/Users/likyam/GoProjects/ygrpc/src/common/config/client.yaml")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %v\n", err)
		return nil
	}

	conf := &Config{}

	// Assign the overloaded configuration to the global
	if err := viper.Unmarshal(conf); err != nil {
		panic("assign config failed.[ERROR]=>" + err.Error())
	}
	fmt.Println(conf)
	return conf

}
