package routers

import (
	"gin-admin/auth"
	"gin-admin/web/index"
	"gin-admin/web/upm"
	"github.com/gin-gonic/gin"
)

// init routers
func InitRouters(r *gin.Engine) {
	// ser user router
	UserRouter(r)
	// index router
	IndexRouter(r)

}

// set index router
func IndexRouter(r *gin.Engine) {
	indexv1 := r.Group("v1/index")
	// 设置权限验证
	indexv1.Use(auth.JwtHandler())
	// 设置router
	{
		indexv1.GET("/", index.Index)
	}
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
	// 注销单独使用handler
	r.POST("/user/logout", auth.JwtHandler(), upm.Logout)
}
