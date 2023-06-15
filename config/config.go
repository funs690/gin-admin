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
			log.Println("config file not found ...")
		} else {
			log.Println("parse error ...")
		}
	}
	return config
}

// init jwt config
func InitJwtConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./conf/")
	config.SetConfigName("jwt")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found ...")
		} else {
			log.Println("parse error ...")
		}
	}
	return config
}
