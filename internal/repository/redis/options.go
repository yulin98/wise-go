package redis

type Options struct {
	Address  string `mapstructure:"address,omitempty"`
	Password string `mapstructure:"password,omitempty"`
	Database int    `mapstructure:"database,omitempty"`
}

func NewRedisRunOptions() *Options {
	return &Options{Address: "redis:6379", Database: 1}
}
