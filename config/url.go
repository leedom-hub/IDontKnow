package config

type Url struct {
	LoginUrl string `mapstructure:"login-url" json:"login-url" yaml:"login-url"`
	LoginOutUrl string `mapstructure:"login-out-url" json:"login-out-url" yaml:"login-out-url"`
	LoginApi string `mapstructure:"login-api" json:"login-api" yaml:"login-api"`
	NatApi string `mapstructure:"nat-api" json:"nat-api" yaml:"nat-api"`
	LocalIp string `mapstructure:"local-ip" json:"local-ip" yaml:"local-ip"`
}
