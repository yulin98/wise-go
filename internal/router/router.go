package router

import (
	"com.wisecharge/central/internal/package/core"
	"com.wisecharge/central/internal/repository/mysql"
	"com.wisecharge/central/internal/repository/redis"
	"com.wisecharge/central/package/errors"
	"go.uber.org/zap"
)

// resource 资源
type resource struct {
	mux   core.Mux
	log   *zap.Logger
	mysql *mysql.GenericMysql
	redis *redis.GenericRedis
}

// Server HTTP 服务器
type Server struct {
	Mux   core.Mux
	Log   *zap.Logger
	Mysql *mysql.GenericMysql
	Redis *redis.GenericRedis
}

// NewHttpServer  HTTP 服务器构造函数
func NewHttpServer(log *zap.Logger) (*Server, error) {
	if log == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.log = log

	// 初始化 db
	dataBase, err := mysql.New(log)
	if err != nil {
		log.Fatal("new db err", zap.Error(err))
	}
	r.mysql = dataBase

	// 初始化 redis
	redisContainer := redis.New()
	r.redis = redisContainer

	// 初始化 http service
	mux, err := core.New(log,
		core.WithEnableCors())

	if err != nil {
		panic(err)
	}

	r.mux = mux

	// 设置 API 路由
	setApiRouter(r)

	s := new(Server)
	s.Mux = r.mux
	s.Mysql = r.mysql
	s.Redis = r.redis
	s.Log = r.log

	return s, nil
}

/*// Start HTTP 服务器
func (hs *Server) Start() {
	go func() {
		hs.Log.Info("start web service. listening port " + hs.Srv.Addr)
		if err := hs.mux.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hs.Log.Error("http service start fail", zap.Error(err))
		}
	}()
}*/
