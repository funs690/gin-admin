package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
)

var db *gorm.DB

// init database
func InitDbLink(config *viper.Viper) *gorm.DB {
	database := config.GetStringMapString("database")
	host := database["host"]
	port, _ := strconv.Atoi(database["port"])
	user := database["username"]
	password := database["password"]
	db_name := database["database"]
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, db_name, port)
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	return db
}
