package user

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 把上面的结构体示例化 这样后可以方便操作
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Migration() {
	// 迁移数据库
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

func (r *Repository) Create(u *User) error {
	// 创建用户
	return r.db.Create(u).Error
}

func (r *Repository) GetByName(name string) (User, error) {
	// 通过用户名查找用户
	var user User
	err := r.db.Where("UserName=?", name).Where("IsDeleted=?", false).First(&user, "UserName=?", name).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin", "admin")
	user.IsAdmin = true
	r.db.Where(User{Username: user.Username}).Attrs(User{Username: user.Username, Password: user.Password}).FirstOrCreate(&user)
	user = NewUser("user", "user", "user")
	r.db.Where(User{Username: user.Username}).Attrs(User{Username: user.Username, Password: user.Password}).FirstOrCreate(&user)

}

func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
