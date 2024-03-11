package user

import (
	"pro05shopping/utils/hash"

	"gorm.io/gorm"
)

// hooks 表示在执行数据库操作前后的一些操作这里就是表示在密码被保存之前给密码后面加一个salt 至于加这个有什么用之前提到过了
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt != "" {
		salt := hash.CreateSalt()
		hashPassword, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return err
		}
		u.Password = hashPassword
		u.Salt = salt
	}
	return nil
}
