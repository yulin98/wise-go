package mysql

type Options struct {
	DSN               string `mapstructure:"dsn,omitempty"`
	MaxIdleConnection int    `mapstructure:"max_idle,omitempty"`
	MaxOpenConnection int    `mapstructure:"max_open,omitempty"`
	MaxLifetime       int    `mapstructure:"max_lifetime,omitempty"`
}

func NewDatabaseOptions() *Options {
	return &Options{
		MaxIdleConnection: 25,
		MaxOpenConnection: 100,
		MaxLifetime:       15,
	}
}
