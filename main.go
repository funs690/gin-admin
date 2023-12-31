package main

import (
	"fmt"
	"gin-admin/config"
	"gin-admin/models"
	"gin-admin/redis"
	"gin-admin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// create gin
	r := gin.Default()
	config.InitApplicationConfig()
	// init database
	models.InitDbLink()
	// init redis
	redis.InitRedisClient()
	// init routes
	routers.InitRouters(r)
	// run server
	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}
