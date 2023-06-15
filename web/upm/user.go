package upm

import (
	http "gin-admin/common"
	"gin-admin/models"
	"gin-admin/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string
	PassWord string
	Name     string
}

// do login
func Login(c *gin.Context) {
	user := User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		http.FailWithMsg(c, 500, "parse parma error")
	}
	userInfo := service.UserService.FindUserByUserName(user.UserName)
	if userInfo == nil || userInfo.PassWord != user.PassWord {
		http.FailWithMsg(c, 401, "login failed")
		return
	}
	// do login
	http.OkWithMsg(c, "login success")
}

func Register(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		http.FailWithMsg(c, 500, "parse parma error")
	}
	service.UserService.Save(user)
	http.OkWithMsg(c, "register success")
}
