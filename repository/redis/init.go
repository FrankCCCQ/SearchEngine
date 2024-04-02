package redis

import (
	"SearchEngine/config"
	"context"
	"fmt"

	logs "SearchEngine/pkg/logger"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func InitRedis() {
	rConfig := config.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		Username: rConfig.RedisUserName,
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDBName,
	})
	if err := redisotel.InstrumentTracing(client); err != nil {
		logs.LogrusObj.Errorf("failed to trace redis, err = %v", err)
	}
	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		logs.LogrusObj.Errorln(err)
		panic(err)
	}
	RedisClient = client

}
