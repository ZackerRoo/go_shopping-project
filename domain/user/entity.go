package user

import (
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Password2 string `gorm:"-"`
	Salt      string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	IsDeleted bool
	IsAdmin   bool
}

// 新建用户实例 新建了一个用户工厂
func NewUser(username, password, password2 string) *User {

	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		IsDeleted: false,
		IsAdmin:   false,
	}
}
