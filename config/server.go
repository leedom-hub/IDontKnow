package config

type Server struct {
	//left 剩余时间
	Counter Counter `mapstructure:"counter" json:"counter" yaml:"counter"`
	//url
	Url Url `mapstructure:"url" json:"url" yaml:"url"`
	//user
	User User `mapstructure:"user" json:"user" yaml:"user"`
	//header
	Header Header `mapstructure:"header" json:"header" yaml:"header"`
	//ZAP
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	//LOGIN FORM
	LoginForm LoginForm `mapstructure:"login-form" json:"login-form" yaml:"login-form"`
	//env
	Env Env `mapstructure:"env" json:"env" yaml:"env"`
}

