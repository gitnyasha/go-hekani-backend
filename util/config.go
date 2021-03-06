package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Driver        string        `mapstructure:"DRIVER"`
	Source        string        `mapstructure:"SOURCE"`
	Server        string        `mapstructure:"SERVER"`
	TokenKey      string        `mapstructure:"TOKEN_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
