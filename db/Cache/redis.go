package cache

import (
	"context"
	"fmt"

	"github.com/aminkhn/portfolio/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	REDIS *redis.Client
	ctx   = context.Background()
)

func RedisConnet() {
	env, _ := config.LoadConfig()
	address := fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: env.RedisPassword, // no password set
		DB:       0,                 // use default DB
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Error().Err(err).Msg("redis connect ping failed")
	} else {
		log.Info().Msg(pong)
		REDIS = client
	}
}

// add health check for redis
func RedisCheck() bool {
	if REDIS != nil {
		_, err := REDIS.Ping(ctx).Result()
		if err != nil {
			log.Error().Err(err).Msg("redis connect ping failed")
			return false
		} else {
			return true
		}
	}
	return false
}

// Close redis connection
func RedisClose() {
	if REDIS != nil {
		err := REDIS.Close()
		if err != nil {
			log.Error().Err(err).Msg("redis close failed")
		}
	}
}
