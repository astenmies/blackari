package utils

import (
	"github.com/spf13/viper"
	"log"
)

func InitViper() {
	// Now we're sure global exists
	viper.SetConfigName("global")
	viper.AddConfigPath("_config")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}
