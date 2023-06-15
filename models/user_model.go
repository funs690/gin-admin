package models

type User struct {
	Id       string `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserName string `gorm:"column:user_name" json:"userName"`
	PassWord string `gorm:"column:pass_word" json:"passWord"`
	Name     string `gorm:"column:name" json:"name"`
}

// 创建用户
func Save(user *User) error {
	result := db.Create(&user)
	return result.Error
}

// 获取单个用户
func FindOne(user *User) *User {
	db.Where(&user).Take(&user)
	return user
}
