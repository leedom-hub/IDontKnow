package config

import "net/http"

type Header struct {
	Accept         string `mapstructure:"accept" json:"accept" yaml:"accept"`
	AcceptEncoding string `mapstructure:"accept-encoding" json:"accept-encoding" yaml:"accept-encoding"`
	AcceptLanguage string `mapstructure:"accept-language" json:"accept-language" yaml:"accept-language"`
	Connection     string `mapstructure:"connection" json:"connection" yaml:"connection"`
	Host           string `mapstructure:"host" json:"host" yaml:"host"`
	Origin         string `mapstructure:"origin" json:"origin" yaml:"origin"`
	Referer        string `mapstructure:"referer" json:"referer" yaml:"referer"`
	UserAgent      string `mapstructure:"user-agent" json:"user-agent" yaml:"user-agent"`
	XCSRFToken     string `mapstructure:"X-CSRFToken" json:"X-CSRFToken" yaml:"X-CSRFToken"`
	XRequestedWith string `mapstructure:"X-Requested-With" json:"X-Requested-With" yaml:"X-Requested-With"`
	Cookie         []*http.Cookie `mapstructure:"cookie" json:"cookie" yaml:"cookie"`
}
