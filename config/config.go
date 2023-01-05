package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config Settings

type Settings struct {
	DB_SERVICE string
}

func init() {
	InitConfig()
}

func InitConfig() {
	viper.AutomaticEnv()
	viper.BindEnv("DB_SERVICE")

	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("Error unmarshalling configuration: %s", err)
	}
}
