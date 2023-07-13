package station

import (
	"com.wisecharge/central/internal/repository/mysql"
	"com.wisecharge/central/internal/repository/redis"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
}

type service struct {
	log   *zap.Logger
	cache *redis.GenericRedis
	mysql *mysql.GenericMysql
}

func New(log *zap.Logger, cache *redis.GenericRedis, mysql *mysql.GenericMysql) Service {
	return &service{
		log:   log,
		cache: cache,
		mysql: mysql,
	}
}
