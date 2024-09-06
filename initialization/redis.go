package initialization

import (
	"github.com/go-redis/redis/v8"
	"github.com/miaoming3/bot/global"
)

func InitContent() {

	global.RedisClient = redis.NewClient(&redis.Options{
		Addr: global.BotConfigure.RedisConfig.Address,
		DB:   0,
	})
}
