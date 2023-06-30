package upm

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-admin/auth"
	"gin-admin/models"
	"gin-admin/redis"
	"github.com/satori/go.uuid"
	"time"
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
func (p *UserMgr) Login(user *models.User) (*auth.Session, error) {
	userInfo := models.FindOne(&models.User{UserName: user.UserName})
	if userInfo.Id == "" {
		return nil, errors.New("user not found")
	}
	if userInfo.PassWord != user.PassWord {
		return nil, errors.New("password is incorrect")
	}
	// generator token
	sessionId := uuid.NewV4().String()

	token, err := auth.GenerateToken(userInfo, sessionId)
	if err != nil || "" == token {
		return nil, errors.New("generator token error")
	}
	// set data
	session := &auth.Session{
		Id:        sessionId,
		UserId:    userInfo.Id,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		CreateAt:  time.Now().Unix(),
		Client:    "",
		Data:      nil,
	}
	data, _ := json.Marshal(session)
	if err := redis.Set(fmt.Sprintf("Authorization:login:session:%s", session.Id), data, 24*time.Hour); err != nil {
		return nil, errors.New("set token cache error")
	}
	return session, err
}

// 获取数据
func (p *UserMgr) Test() *[]models.User {
	users := models.FindList(&models.User{Name: "张三"})
	for _, user := range *users {
		fmt.Print(user.UserName)
	}
	return users
}

// 用户推出登录
func (p *UserMgr) Logout(sessionId string) error {
	return redis.Delete(fmt.Sprintf("Authorization:login:session:%s", sessionId))
}
