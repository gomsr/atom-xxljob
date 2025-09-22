package configx

import "time"

const (
	defaultAdminAddr = "http://localhost:8080/xxl-job-admin/"
	defaultAppName   = "go-executor"
	defaultPort      = 8081
	defaultTimeout   = 5 * time.Second
	defaultBeatTime  = 20 * time.Second
)

type Option func(*ClientConfig)

func NewClientOptions(opts ...Option) ClientConfig {
	options := ClientConfig{
		AdminAddr:   []string{defaultAdminAddr},
		AccessToken: "",
		AppName:     defaultAppName,
		Port:        defaultPort,
		Timeout:     defaultTimeout,
		BeatTime:    defaultBeatTime,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// WithAdminAddress xxl admin address
func WithAdminAddress(addrs ...string) Option {
	return func(o *ClientConfig) {
		o.AdminAddr = addrs
	}
}

// WithAccessToken xxl admin accessToke
func WithAccessToken(token string) Option {
	return func(o *ClientConfig) {
		o.AccessToken = token
	}
}

// WithAppName app name
func WithAppName(appName string) Option {
	return func(o *ClientConfig) {
		o.AppName = appName
	}
}

// WithClientPort xxl client port
func WithClientPort(port int) Option {
	return func(o *ClientConfig) {
		o.Port = port
	}
}

// WithAdminTimeout xxl admin request timeout
func WithAdminTimeout(timeout time.Duration) Option {
	return func(o *ClientConfig) {
		o.Timeout = timeout
	}
}

// WithBeatTime xxl admin renew time
func WithBeatTime(beatTime time.Duration) Option {
	return func(o *ClientConfig) {
		o.BeatTime = beatTime
	}
}

func WithLogLevel(level int) Option {
	return func(o *ClientConfig) {
		o.LogLevel = level
	}
}
