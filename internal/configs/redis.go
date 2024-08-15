package configs

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ctx = context.Background()

func NewRedis(viper *viper.Viper, log *zap.Logger) *redis.Client {

	redisDb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	pong, err := redisDb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("failed to connect redis", zap.Error(err))
	}
	log.Info("connected to redis", zap.String("ping", pong))
	return redisDb
}
