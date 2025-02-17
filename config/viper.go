package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AllowEmptyEnv(false)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("No .env file found, using environment variables.")
	}

	return v
}
