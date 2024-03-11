package category

import (
	"gorm.io/gorm"
)

// 商品分类结构体，对应数据库表
type Category struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Desc     string
	IsActive bool
}

// 新建商品分类
func NewCategory(name string, desc string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}
}
