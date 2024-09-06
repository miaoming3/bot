package config

type BotConfig struct {
	AppId       uint64      `mapstructure:"AppID" json:"AppID"`
	Token       string      `mapstructure:"Token" json:"Token"`
	AppSecret   string      `mapstructure:"AppSecret" json:"AppSecret"`
	RedisConfig NosqlConfig `mapstructure:"redis" json:"redis"`
}

type NosqlConfig struct {
	Address string `mapstructure:"address" json:"address"`
}
