package initialization

import (
	"fmt"
	"github.com/miaoming3/bot/global"
	"github.com/spf13/viper"
)

func LoadConfig(file string) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := viper.Unmarshal(&global.BotConfigure); err != nil {
		panic(err)
	}
	fmt.Println(global.BotConfigure)
}
