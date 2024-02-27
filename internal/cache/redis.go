package cache

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/nvlhnn/gommerce/internal/config"
)

func CreateClient(cfg config.Config ) (*redis.Client, error) {
	log.Println(cfg.RedisAddr, cfg.RedisPort, cfg.RedisPass, cfg.RedisDB)

	opt := &redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.RedisAddr, cfg.RedisPort),
		DB:       cfg.RedisDB,
	}

	if cfg.RedisPass != "" {
		opt.Password = cfg.RedisPass
	}

	rdb := redis.NewClient(opt)

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, err
	}


	return rdb, nil
}