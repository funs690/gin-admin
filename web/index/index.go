package index

import (
	"gin-admin/common"
	service "gin-admin/service"
	"github.com/gin-gonic/gin"
	"log"
)

// index test
func Index(c *gin.Context) {
	// test uid
	log.Print(c.GetString("uid"))
	users := service.UserService.Test()
	common.OkWithData(c, users)
}
