package config

type Env struct {
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
	Proxy bool `mapstructure:"proxy" json:"proxy" yaml:"proxy"`
	OutPutCount bool `mapstructure:"out-put-count" json:"out-put-count" yaml:"out-put-count"`
}

