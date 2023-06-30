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

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

// set config
var (
	Server   *ServerConfig
	Database *DataBaseConfig
	Jwt      *JwtConfig
	Redis    *RedisConfig
	Viper    *viper.Viper
)

// init config
func InitApplicationConfig() {
	Viper = viper.New()
	Viper.AddConfigPath("./conf/")
	Viper.SetConfigName("application")
	Viper.SetConfigType("yaml")
	if err := Viper.ReadInConfig(); err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
	//init server config
	InitServerConfig(Viper)
	//init db config
	InitDataBaseConfig(Viper)
	//init jwt config
	InitJwtConfig(Viper)
	//init redis config
	InitRedisConfig(Viper)
}

// init redis config
func InitRedisConfig(viper *viper.Viper) {
	Redis = new(RedisConfig)
	config := viper.GetStringMapString("redis")
	Redis.Host = config["host"]
	if port, err := strconv.Atoi(config["port"]); err != nil {
		// set default
		Redis.Port = 6379
	} else {
		Redis.Port = port
	}
	Redis.Password = config["password"]
	if database, err := strconv.Atoi(config["database"]); err != nil {
		// set default
		Redis.Database = 0
	} else {
		Redis.Database = database
	}
}

// set jwt info
func InitJwtConfig(viper *viper.Viper) {
	Jwt = new(JwtConfig)
	config := viper.GetStringMapString("jwt")
	Jwt.Sign = config["sign"]
	Jwt.Secret = config["secret"]
}

// set database info
func InitDataBaseConfig(viper *viper.Viper) {
	Database = new(DataBaseConfig)
	config := viper.GetStringMapString("database")
	Database.Host = config["host"]
	if port, err := strconv.Atoi(config["port"]); err != nil {
		// set default
		Database.Port = 0
	} else {
		Database.Port = port
	}
	Database.Username = config["username"]
	Database.Password = config["password"]
	Database.Database = config["database"]
}

// set server info
func InitServerConfig(viper *viper.Viper) {
	Server = new(ServerConfig)
	config := viper.GetStringMapString("server")
	if port, err := strconv.Atoi(config["port"]); err != nil {
		// set default
		Server.Port = 8080
	} else {
		Server.Port = port
	}
	if addr := config["addr"]; addr == "" {
		Server.Addr = "0.0.0.0"
	} else {
		Server.Addr = addr
	}
}
