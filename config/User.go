package config

type User struct {
	UserName string `mapstructure:"user-name" json:"user-name" yaml:"user-name"`
	Pwd string `mapstructure:"pwd" json:"pwd" yaml:"pwd"`
}

