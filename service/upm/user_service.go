package upm

import (
	"errors"
	"fmt"
	"gin-admin/auth"
	"gin-admin/models"
)

// define
type UserService interface {
	// 用户保存
	Save(user *models.User)
	// 根据用户名查找的用户
	FindUserByUserName(userName string) *models.User
}

// define service
type UserMgr struct {
}

// user save
func (p *UserMgr) Register(user *models.User) error {
	return models.Save(user)
}

// user login
func (p *UserMgr) Login(user *models.User) (string, error) {
	userInfo := models.FindOne(&models.User{UserName: user.UserName})
	if userInfo.Id == "" {
		return "", errors.New("user not found")
	}
	if userInfo.PassWord != user.PassWord {
		return "", errors.New("password is incorrect")
	}
	// generator token
	return auth.GenerateToken(userInfo)
}

// 获取数据
func (p *UserMgr) Test() *[]models.User {
	users := models.FindList(&models.User{Name: "张三"})
	for _, user := range *users {
		fmt.Print(user.UserName)
	}
	return users
}
