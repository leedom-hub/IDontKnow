package config

type LoginForm struct {
	UserName string `mapstructure:"user-name" json:"user-name" yaml:"user-name"`
	Pwd string `mapstructure:"pwd" json:"pwd" yaml:"pwd"`
	RequestUrl string `mapstructure:"request-url" json:"request-url" yaml:"request-url"`
	OsName string `mapstructure:"os-name" json:"os-name" yaml:"os-name"`
	BrowserName string `mapstructure:"browser-name" json:"browser-name" yaml:"browser-name"`
	ForceChange string `mapstructure:"force-change" json:"force-change" yaml:"force-change"`
}
