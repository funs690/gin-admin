package routers

import (
	"gin-admin/web/upm"
	"github.com/gin-gonic/gin"
)

// init routers
func InitRouters(r *gin.Engine) {
	// ser user router
	UserRouter(r)
}

// user router
func UserRouter(r *gin.Engine) {
	// set router group
	user := r.Group("user")
	{
		// user register
		user.POST("/register", upm.Register)
		// add user router
		user.POST("/login", upm.Login)
	}
}
