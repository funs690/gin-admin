package main

import (
	"fmt"
	"gin-admin/config"
	models "gin-admin/models"
	"gin-admin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// create gin
	r := gin.Default()
	config := config.InitApplicationConfig()
	// init database
	models.InitDbLink(config)
	// init routes
	routers.InitRouters(r)
	// run server
	r.Run(fmt.Sprintf(":%d", config.GetInt("port")))
}
