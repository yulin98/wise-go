package core

import (
	"com.wisecharge/central/internal/code"
	"com.wisecharge/central/internal/proposal"
	"com.wisecharge/central/package/env"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cors "github.com/rs/cors/wrapper/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"runtime/debug"
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

type Option func(*option)

type option struct {
	disablePProf      bool
	disableSwagger    bool
	disablePrometheus bool
	enableCors        bool
	enableRate        bool
	enableOpenBrowser string
}

// WithDisablePProf 禁用 pprof
func WithDisablePProf() Option {
	return func(opt *option) {
		opt.disablePProf = true
	}
}

// WithDisableSwagger 禁用 swagger
func WithDisableSwagger() Option {
	return func(opt *option) {
		opt.disableSwagger = true
	}
}

// WithDisablePrometheus 禁用prometheus
func WithDisablePrometheus() Option {
	return func(opt *option) {
		opt.disablePrometheus = true
	}
}

// WithEnableOpenBrowser 启动后在浏览器中打开 uri
func WithEnableOpenBrowser(uri string) Option {
	return func(opt *option) {
		opt.enableOpenBrowser = uri
	}
}

// WithEnableCors 设置支持跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// WithEnableRate 设置支持限流
func WithEnableRate() Option {
	return func(opt *option) {
		opt.enableRate = true
	}
}

// DisableRecordMetrics 禁止记录指标
func DisableRecordMetrics(ctx Context) {
	ctx.disableRecordMetrics()
}

// AliasForRecordMetrics 对请求路径起个别名，用于记录指标。
// 如：Get /user/:username 这样的路径，因为 username 会有非常多的情况，这样记录指标非常不友好。
func AliasForRecordMetrics(path string) HandlerFunc {
	return func(ctx Context) {
		ctx.setAlias(path)
	}
}

// WrapAuthHandler 用来处理 Auth 的入口
func WrapAuthHandler(handler func(Context) (sessionUserInfo proposal.SessionUserInfo, err BusinessError)) HandlerFunc {
	return func(ctx Context) {
		sessionUserInfo, err := handler(ctx)
		if err != nil {
			ctx.AbortWithError(err)
			return
		}

		ctx.setSessionUserInfo(sessionUserInfo)
	}
}

// RouterGroup 包装gin的RouterGroup
type RouterGroup interface {
	Group(string, ...HandlerFunc) RouterGroup
	IRoutes
}

var _ IRoutes = (*router)(nil)

// IRoutes 包装gin的IRoutes
type IRoutes interface {
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

type router struct {
	group *gin.RouterGroup
}

func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

func (r *router) Any(relativePath string, handlers ...HandlerFunc) {
	r.group.Any(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.group.DELETE(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	r.group.OPTIONS(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	r.group.HEAD(relativePath, wrapHandlers(handlers...)...)
}

func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)

			handler(ctx)
		}
	}

	return funcs
}

var _ Mux = (*mux)(nil)

// Mux http mux
type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func New(logger *zap.Logger, options ...Option) (Mux, error) {
	//仅关闭 GIN-debug 日志
	gin.DefaultWriter = io.Discard

	mux := &mux{
		engine: gin.New(),
	}
	fmt.Println(fmt.Sprintf("%s", _UI))

	// option
	opt := new(option)
	for _, f := range options {
		f(opt)
	}
	if !opt.disablePProf {
		if !env.Active().IsPro() {
			pprof.Register(mux.engine) // register pprof to gin
		}
	}
	if !opt.disablePrometheus {
		mux.engine.GET("/metrics", gin.WrapH(promhttp.Handler())) // register prometheus
	}

	if opt.enableCors {
		mux.engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	// recover两次，防止处理时发生panic，
	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
			// handler error
			return
		}()

		gin.LoggerWithConfig(
			gin.LoggerConfig{SkipPaths: []string{"/system/health", "/metrics"}},
		)
		ctx.Next()
	})

	mux.engine.Use(func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}
		context := newContext(ctx)
		defer releaseContext(context)

		context.init()
		context.setLogger(logger)

		defer func() {
			var (
				response interface{}
			)
			// region 错误处理 todo ...
			if err := recover(); err != nil {
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
				context.AbortWithError(
					Error(http.StatusInternalServerError,
						code.ServerError,
						err.(error).Error()),
				)
			}
			if ctx.IsAborted() {
				if err := context.abortError(); err != nil {
					response = &code.FinalResponse{
						Code:    err.BusinessCode(),
						Message: err.Message(),
						Body:    nil,
					}
					ctx.JSON(http.StatusOK, response)
				}
			}

			// region 正确返回
			response = context.getPayload()
			if response != nil {
				ctx.JSON(http.StatusOK, response)
			}
		}()

		ctx.Next()
	})

	system := mux.Group("/system")
	{
		// 健康检查
		system.GET("/health", func(ctx Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Host(),
				Status:      "ok",
			}
			ctx.Payload(resp)
		})
	}

	return mux, nil
}
