package order

import (
	"pro05shopping/domain/cart"
	"pro05shopping/domain/product"

	"gorm.io/gorm"
)

// BeforeCreate 在创建订单之前被调用
// 它查找用户的购物车并删除它。
func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	// 查找当前用户的购物车
	var currentCart cart.Cart // 存储当前用户的购物车的变量

	// 通过用户ID在数据库中查找购物车
	if err := tx.Where("UserID = ?", order.UserID).First(&currentCart).Error; err != nil {
		return err
	}

	// 删除购物车中的所有商品项
	// 通过购物车ID在数据库中删除购物车中的商品项
	if err := tx.Where("CartID = ?", currentCart.ID).Unscoped().Delete(&cart.Item{}).Error; err != nil {
		return err
	}

	// 删除当前购物车
	// 从数据库中删除当前用户的购物车
	if err := tx.Unscoped().Delete(&currentCart).Error; err != nil {
		return err
	}

	return nil
}

// BeforeSave 在保存订单项之前被调用
// 它更新产品库存。
func (orderedItem *OrderedItem) BeforeSave(tx *gorm.DB) (err error) {
	// 查找当前订单项的商品
	var currentProduct product.Product // 存储当前订单项的商品的变量
	var currentOrderedItem OrderedItem // 存储当前订单项的变量

	// 通过商品ID在数据库中查找当前订单项的商品
	if err := tx.Where("ID = ?", orderedItem.ProductID).First(&currentProduct).Error; err != nil {
		return err
	}

	// 查找当前订单项的数量
	reservedStockCount := 0 // 存储当前订单项的数量的变量
	if err := tx.Where("ID = ?", orderedItem.ID).First(&currentOrderedItem).Error; err == nil {
		reservedStockCount = currentOrderedItem.Count
	}

	// 计算新的库存数量
	newStockCount := currentProduct.StockCount + reservedStockCount - orderedItem.Count
	if newStockCount < 0 {
		return ErrNotEnoughStock
	}

	// 更新产品库存
	// 更新当前商品的库存数量到数据库
	if err := tx.Model(&currentProduct).Update("StockCount", newStockCount).Error; err != nil {
		return err
	}

	// 如果订单项的数量为0，则从数据库中删除该订单项
	if orderedItem.Count == 0 {
		err := tx.Unscoped().Delete(currentOrderedItem).Error
		return err
	}

	return nil
}

// BeforeUpdate 在更新订单之前被调用
// 如果订单被取消，金额将返回到产品库存中。
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	// 如果订单被取消，则将金额返回到产品库存中
	if order.IsCanceled {
		var orderedItems []OrderedItem // 存储订单项的变量
		// 通过订单ID在数据库中查找所有订单项
		if err := tx.Where("OrderID = ?", order.ID).Find(&orderedItems).Error; err != nil {
			return err
		}
		for _, item := range orderedItems {
			var currentProduct product.Product // 存储当前订单项的商品的变量
			// 通过订单项的商品ID在数据库中查找当前订单项的商品
			if err := tx.Where("ID = ?", item.ProductID).First(&currentProduct).Error; err != nil {
				return err
			}

			// 计算新的库存数量
			newStockCount := currentProduct.StockCount + item.Count
			// 将新的库存数量更新到数据库中
			if err := tx.Model(&currentProduct).Update("StockCount", newStockCount).Error; err != nil {
				return err
			}

			// 将订单项的IsCanceled字段设置为true，表示订单项已取消
			if err := tx.Model(&item).Update("IsCanceled", true).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
