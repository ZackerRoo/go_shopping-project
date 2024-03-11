package user

import (
	"errors"
)

// 用户自定义错误

var (
	ErrUserExistWithName = errors.New("用户名已经存在")
	ErrUserNotFound      = errors.New("用户未找到")

	ErrMismatchedPasswords = errors.New("密码不匹配")
	ErrInvalidUsername     = errors.New("无效用户名")
	ErrInvalidPassword     = errors.New("无效密码")
)
