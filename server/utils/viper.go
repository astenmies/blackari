package utils

import (
	"log"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("global")
	viper.AddConfigPath("_config")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}
