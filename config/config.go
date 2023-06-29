package config

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// server config
type ServerConfig struct {
	Addr string
	Port int
}

// database config
type DataBaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// jwt config
type JwtConfig struct {
	Secret string
	Sign   string
}

// set config
var (
	Server   *ServerConfig
	Database *DataBaseConfig
	Jwt      *JwtConfig
)

// init config
func InitApplicationConfig() {
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
	//init server config
	InitServerConfig(config)
	//init db config
	InitDataBaseConfig(config)
	//init jwt config
	InitJwtConfig(config)
}

// set jwt info
func InitJwtConfig(config *viper.Viper) {
	Jwt = new(JwtConfig)
	jwtConfig := config.GetStringMapString("jwt")
	Jwt.Sign = jwtConfig["sign"]
	Jwt.Secret = jwtConfig["secret"]
}

// set database info
func InitDataBaseConfig(config *viper.Viper) {
	Database = new(DataBaseConfig)
	databaseConfig := config.GetStringMapString("database")
	Database.Host = databaseConfig["host"]
	if port, err := strconv.Atoi(databaseConfig["port"]); err != nil {
		// set default
		Database.Port = 0
	} else {
		Database.Port = port
	}
	Database.Username = databaseConfig["username"]
	Database.Password = databaseConfig["password"]
	Database.Database = databaseConfig["database"]
}

// set server info
func InitServerConfig(config *viper.Viper) {
	Server = new(ServerConfig)
	serverConfig := config.GetStringMapString("server")
	if port, err := strconv.Atoi(serverConfig["port"]); err != nil {
		// set default
		Server.Port = 8080
	} else {
		Server.Port = port
	}
	if addr := serverConfig["addr"]; addr == "" {
		Server.Addr = "0.0.0.0"
	} else {
		Server.Addr = addr
	}
}
