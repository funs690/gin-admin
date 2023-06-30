package upm

import (
	http "gin-admin/common"
	"gin-admin/models"
	"gin-admin/service"
	"github.com/gin-gonic/gin"
	"log"
)

// do login
func Login(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		http.FailWithMsg(c, 500, "parse parma error")
	}
	session, err := service.UserService.Login(user)
	if err != nil {
		http.FailWithMsg(c, 500, err.Error())
		return
	}
	// do login
	http.OkWithData(c, session)
}

// do register
func Register(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		http.FailWithMsg(c, 500, "parse parma error")
	}
	if err := service.UserService.Register(user); err != nil {
		log.Println(err)
		http.FailWithMsg(c, 500, "user register failed")
	}
	http.OkWithMsg(c, "register success")
}

// do logout
func Logout(c *gin.Context) {
	sessionId := c.GetString("sessionId")
	err := service.UserService.Logout(sessionId)
	if err != nil {
		http.FailWithMsg(c, 500, err.Error())
		return
	}
	http.OkWithMsg(c, "login out success")
}
