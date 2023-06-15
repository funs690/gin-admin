package upm

import (
	http "gin-admin/common"
	"gin-admin/models"
	"gin-admin/service"
	"github.com/gin-gonic/gin"
	"log"
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
	if userInfo == nil {
		http.FailWithMsg(c, 401, "user not found")
		return
	}
	if userInfo.PassWord != user.PassWord {
		http.FailWithMsg(c, 401, "pass world is incorrect")
		return
	}
	// do login
	http.OkWithMsg(c, "login success")
}

// do register
func Register(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		http.FailWithMsg(c, 500, "parse parma error")
	}
	if err := service.UserService.Save(user); err != nil {
		log.Println(err)
		http.FailWithMsg(c, 500, "create user failed")
	}
	http.OkWithMsg(c, "register success")
}
