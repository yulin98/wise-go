package server

import (
	"com.wisecharge/central/configs"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"time"
)

const _UI = `
               /$$                                       /$$                                              
              |__/                                      | $$                                              
 /$$  /$$  /$$ /$$  /$$$$$$$  /$$$$$$           /$$$$$$$| $$$$$$$   /$$$$$$   /$$$$$$   /$$$$$$   /$$$$$$ 
| $$ | $$ | $$| $$ /$$_____/ /$$__  $$ /$$$$$$ /$$_____/| $$__  $$ |____  $$ /$$__  $$ /$$__  $$ /$$__  $$
| $$ | $$ | $$| $$|  $$$$$$ | $$$$$$$$|______/| $$      | $$  \ $$  /$$$$$$$| $$  \__/| $$  \ $$| $$$$$$$$
| $$ | $$ | $$| $$ \____  $$| $$_____/        | $$      | $$  | $$ /$$__  $$| $$      | $$  | $$| $$_____/
|  $$$$$/$$$$/| $$ /$$$$$$$/|  $$$$$$$        |  $$$$$$$| $$  | $$|  $$$$$$$| $$      |  $$$$$$$|  $$$$$$$
 \_____/\___/ |__/|_______/  \_______/         \_______/|__/  |__/ \_______/|__/       \____  $$ \_______/
                                                                                       /$$  \ $$          
                                                                                      |  $$$$$$/          
                                                                                       \______/            
`

/*func logFormatter(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[wise-central] %v |%3d| %13v |%-7s %s\n%s",
		param.TimeStamp.Format("2006/01/02 15:04:05.999999"),
		param.StatusCode,
		param.Latency,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}*/

func New() *http.Server {
	cfg := configs.Get().Server
	options := Options{
		BindAddress:           cfg.BindAddress,
		InsecurePort:          cfg.InsecurePort,
		ShutdownDelayDuration: cfg.ShutdownDelayDuration,
	}

	//仅关闭 GIN-debug 日志
	gin.DefaultWriter = io.Discard
	router := gin.New()
	router.Use(gin.Recovery(), gin.LoggerWithConfig(
		gin.LoggerConfig{SkipPaths: []string{"/health", "/metrics"}},
	))
	pprof.Register(router, "debug/pprof")

	// skipPaths /metrics、 /health
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "OK",
			"Timestamp": time.Now(),
			"host":      ctx.Request.Host,
		})
	})

	// 创建一个 http.Server 对象
	server := &http.Server{
		Addr:    options.InsecurePort,
		Handler: router,
	}
	fmt.Println(fmt.Sprintf("%s", _UI))
	return server
}
