package connector

import (
	"com.wisecharge/central/internal/package/core"
	"com.wisecharge/central/internal/repository/mysql"
	"com.wisecharge/central/internal/repository/redis"
	"go.uber.org/zap"
)

type Handler interface {
	CreateConnector() core.HandlerFunc

	DeleteConnector() core.HandlerFunc

	UpdateConnector() core.HandlerFunc

	QueryOneConnector() core.HandlerFunc

	QueryPageConnector() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  *redis.GenericRedis
	mysql  *mysql.GenericMysql
}

func New(log *zap.Logger, cache *redis.GenericRedis, db *mysql.GenericMysql) Handler {
	return &handler{
		logger: log,
		cache:  cache,
		mysql:  db,
	}
}
