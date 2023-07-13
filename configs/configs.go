package configs

import (
	"bytes"
	"com.wisecharge/central/package/env"
	"com.wisecharge/central/package/file"
	_ "embed"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"time"
)

var config = new(Config)

type Config struct {
	// Server 服务配置
	Server struct {
		ApplicationName string

		// server bind address
		BindAddress  string `mapstructure:"address,omitempty"`
		InsecurePort string `mapstructure:"port,omitempty"`

		RequestTimeout      time.Duration
		MaxRequestBodyBytes int64

		ShutdownDelayDuration time.Duration
	} `mapstructure:"server,omitempty"`

	// Logger 日志配置
	Logger struct {
		DebugMode bool `mapstructure:"debug,omitempty"`
	} `mapstructure:"log,omitempty"`

	// MySQL 数据库配置
	MySQL struct {
		DSN               string `mapstructure:"dsn,omitempty"`
		MaxIdleConnection int    `mapstructure:"max_idle,omitempty"`
		MaxOpenConnection int    `mapstructure:"max_open,omitempty"`
		MaxLifetime       int    `mapstructure:"max_lifetime,omitempty"`
	} `mapstructure:"database,omitempty"`

	// Redis Redis配置
	Redis struct {
		Address  string `mapstructure:"address,omitempty"`
		Password string `mapstructure:"password,omitempty"`
		Database int    `mapstructure:"database,omitempty"`
	} `mapstructure:"redis,omitempty"`
}

var (
	//go:embed dev_configs.properties
	devConfigs []byte

	//go:embed fat_configs.properties
	fatConfigs []byte

	//go:embed pro_configs.properties
	uatConfigs []byte

	//go:embed uat_configs.properties
	proConfigs []byte
)

func init() {
	var r io.Reader

	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfigs)
	case "fat":
		r = bytes.NewReader(fatConfigs)
	case "uat":
		r = bytes.NewReader(uatConfigs)
	case "pro":
		r = bytes.NewReader(proConfigs)
	default:
		r = bytes.NewReader(fatConfigs)
	}

	viper.SetConfigType("properties")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.AddConfigPath("./configs")

	configFile := "./configs/" + env.Active().Value() + "_configs.properties"
	_, ok := file.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}
