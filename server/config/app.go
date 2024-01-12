package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	Addr    string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
