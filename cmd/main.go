package main

import (
	"com.wisecharge/central/configs"
	"com.wisecharge/central/internal/log"
	"com.wisecharge/central/internal/router"
	"com.wisecharge/central/package/shutdown"
	"context"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	//初始化 logger
	logger, err := log.New()
	if err != nil {
		panic(err)
	}

	// 关闭 log , 确保日志消息在程序结束时被立即写入输出，而不会被丢失或延迟写入。
	defer func() {
		_ = logger.Sync()
	}()

	// 初始化 HTTP 服务
	hs, err := router.NewHttpServer(logger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.Get().Server.InsecurePort,
		Handler: hs.Mux,
	}

	go func() {
		logger.Info("http server startup ,listening port : " + configs.Get().Server.InsecurePort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(

		// 关闭 web server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				hs.Log.Error("http service shutdown", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if hs.Mysql != nil {
				if err := hs.Mysql.Shutdown(); err != nil {
					hs.Log.Error("db close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if hs.Redis != nil {
				if err := hs.Redis.Shutdown(); err != nil {
					hs.Log.Error("db redis err", zap.Error(err))
				}
			}
		},
	)
}
