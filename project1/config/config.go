package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	GiphyAPIKEY string `mapstructure:"giphy_api_key"`
}

func New(name string) *Config {
	env := os.Getenv("ENV")
	filename := name
	if env != "" {
		filename = filename + "." + env + ".yaml"
	}
	viper.AddConfigPath("config")
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("reading config err: %s", err.Error()))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("unmarshaling config values err: %s", err.Error()))
	}
	return &config
}
