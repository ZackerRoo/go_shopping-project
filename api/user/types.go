package user

// 创建用户请求结构体
type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

// 创建用户响应
type CreateUserResponse struct {
	Username string `json:"username"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录响应
type LoginResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	Token    string `json:"token"`
}
