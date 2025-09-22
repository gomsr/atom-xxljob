package configx

import "time"

type ClientConfig struct {
	// xxl admin 地址
	AdminAddr []string `mapstructure:"admin-addr" json:"admin-addr" yaml:"admin-addr"`

	//token
	AccessToken string `mapstructure:"access-token" json:"access-token" yaml:"access-token"`

	//执行期名
	AppName string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`

	//执行器端口
	Port int `mapstructure:"port" json:"port" yaml:"port"`

	//请求admin超时时间
	Timeout time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`

	//执行器续约时间（超过30秒不续约admin会移除执行器，请设置到30秒以内）
	BeatTime time.Duration `mapstructure:"beat-time" json:"beat-time" yaml:"beat-time"`

	LogLevel int `mapstructure:"log-level" json:"log-level" yaml:"log-level"`
}
