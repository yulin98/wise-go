package server

import "time"

type Options struct {
	ApplicationName string

	// server bind address
	BindAddress  string `mapstructure:"address,omitempty"`
	InsecurePort string `mapstructure:"port,omitempty"`

	RequestTimeout      time.Duration
	MaxRequestBodyBytes int64

	ShutdownDelayDuration time.Duration
}
