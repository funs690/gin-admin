package config

import (
	"github.com/spf13/viper"
	"log"
)

// init config
func InitApplicationConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./conf/")
	config.SetConfigName("application")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found ...")
		} else {
			log.Fatal("parse error ...")
		}
	}
	return config
}
