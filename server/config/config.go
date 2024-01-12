package config

type Server struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Harbor Harbor `mapstructure:"harbor" json:"harbor" yaml:"harbor"`
}
