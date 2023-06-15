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
	appConfig := config.InitApplicationConfig()
	// init database
	models.InitDbLink(appConfig)
	// init routes
	routers.InitRouters(r)
	// run server
	r.Run(fmt.Sprintf(":%d", appConfig.GetInt("port")))
}
