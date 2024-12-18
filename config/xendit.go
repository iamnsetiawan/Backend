package config

import (
	"github.com/spf13/viper"
	xendit "github.com/xendit/xendit-go/v6"
)

func NewXendit(viper *viper.Viper) *xendit.APIClient {
	return xendit.NewClient(viper.GetString("XENDIT_API_KEY"))
}