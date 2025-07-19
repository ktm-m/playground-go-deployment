package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App   App
	HTTP  HTTP
	DB    DB
	Redis Redis
}

type App struct {
	Name string
}

type HTTP struct {
	Port string
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Dialect  string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic("[Config] cannot read config: " + err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic("[Config] cannot unmarshal config: " + err.Error())
	}

	return &config
}
