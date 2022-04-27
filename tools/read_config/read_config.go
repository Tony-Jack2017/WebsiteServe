package read_config

import (
	"Serve/tools/color_print"
	"github.com/spf13/viper"
)

var config Config

func ReadConfig() Config {
	viper.SetConfigName("settings")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		color_print.ColorPrintln("读取配置文件错误！Error:",color_print.Fail(err))
		panic(err)
	}
	viper.Unmarshal(&config)
	return config
}