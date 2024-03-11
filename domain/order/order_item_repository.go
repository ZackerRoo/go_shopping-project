package order

import (
	"log"

	"gorm.io/gorm"
)

type OrderedItemRepository struct {
	db *gorm.DB
}

// 实例化
func NewOrderedItemRepository(db *gorm.DB) *OrderedItemRepository {
	return &OrderedItemRepository{
		db: db,
	}
}

// 创建表
func (r *OrderedItemRepository) Migration() {
	err := r.db.AutoMigrate(&OrderedItem{})
	if err != nil {
		log.Print(err)
	}
}

// 更新
func (r *OrderedItemRepository) Update(item OrderedItem) error {
	result := r.db.Save(&item)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 创建订单item
func (r *OrderedItemRepository) Create(ci *OrderedItem) error {
	result := r.db.Create(ci)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
