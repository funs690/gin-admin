package index

import (
	"gin-admin/common"
	"github.com/gin-gonic/gin"
)

// index test
func Index(c *gin.Context) {
	common.OkWithMsg(c, "this is index!")
}
