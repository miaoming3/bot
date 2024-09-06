package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/miaoming3/bot/config"
)

var (
	BotConfigure *config.BotConfig
	RedisClient  *redis.Client
)
