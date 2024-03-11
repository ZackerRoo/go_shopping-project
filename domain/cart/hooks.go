package cart

import (
	"gorm.io/gorm"
)

// 如果计数为零，则删除商品 这个hook就是说在更新之后时候如果商品数量为0就删除商品
func (item *Item) AfterUpdate(tx *gorm.DB) (err error) {

	if item.Count <= 0 {
		return tx.Unscoped().Delete(&item).Error
	}
	return
}
