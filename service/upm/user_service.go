package upm

import (
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
type UserSvc struct {
}

// user save
func (p *UserSvc) Save(user *models.User) error {
	return models.Save(user)
}

// 查询用户
func (p *UserSvc) FindUserByUserName(userName string) *models.User {
	user := new(models.User)
	user.UserName = userName
	return models.FindOne(user)
}
