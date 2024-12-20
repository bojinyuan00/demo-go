package global

import (
	"demo-go/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config  config.Config
	GormDB  *gorm.DB
	RedisDb *redis.Client
)
