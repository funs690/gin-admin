package index

import (
	"gin-admin/common"
	"gin-admin/service"
	"github.com/gin-gonic/gin"
	"log"
)

// index test
func Index(c *gin.Context) {
	// test uid
	log.Print(c.GetString("uid"))
	service.UserService.Test()
	common.OkWithMsg(c, "this is index!")
}
