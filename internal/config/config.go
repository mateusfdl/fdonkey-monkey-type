package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config interface {
	Get(field string) interface{}
}

func init() {
	viper.SetConfigName("theme")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/fdonkey")
}

func LoadConfig() Config {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper.GetViper()
}
