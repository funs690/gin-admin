package index

import (
	"gin-admin/common"
	"github.com/gin-gonic/gin"
	"log"
)

// index test
func Index(c *gin.Context) {
	// test uid
	log.Print(c.GetString("uid"))
	common.OkWithMsg(c, "this is index!")
}
