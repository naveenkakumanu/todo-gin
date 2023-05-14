package config

import "github.com/spf13/viper"

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{}
}

func (conf *Config) Config() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	conf.Port = viper.GetString("PORT")
	return conf
}
