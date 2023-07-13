package redis

import (
	"com.wisecharge/central/configs"
	"github.com/go-redis/redis/v8"
)

type GenericRedis struct {
	delegate *redis.Client
}

func New() *GenericRedis {
	cfg := configs.Get().Redis
	return &GenericRedis{
		delegate: redis.NewClient(&redis.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       cfg.Database,
		}),
	}
}

func (g GenericRedis) Shutdown() error {
	return g.delegate.Close()
}

func (g GenericRedis) Delegate() *redis.Client {
	return g.delegate
}
