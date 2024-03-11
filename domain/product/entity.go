package product

import (
	"pro05shopping/domain/category"

	"gorm.io/gorm"
)

// 商品结构体
type Product struct {
	gorm.Model
	Name       string
	SKU        string
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint              // 分类id
	Category   category.Category `json:"-"` // 分类
	IsDeleted  bool
}

// 商品结构体实例
func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}
