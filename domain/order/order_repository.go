package order

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewOrderRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 创建表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Order{})
	if err != nil {
		log.Print(err)
	}
}

// 根据订单id查找
func (r *Repository) FindByOrderID(oid uint) (*Order, error) {
	var currentOrder *Order
	if err := r.db.Where("IsCanceled = ?", false).Where("ID", oid).First(&currentOrder).Error; err != nil {
		return nil, err
	}
	return currentOrder, nil

}

// 更新订单
func (r *Repository) Update(newOrder Order) error {
	result := r.db.Save(&newOrder)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 创建订单
func (r *Repository) Create(ci *Order) error {
	result := r.db.Create(ci)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 获得所有订单
func (r *Repository) GetAll(pageIndex, pageSize int, uid uint) ([]Order, int) {
	var orders []Order
	var count int64

	r.db.Where("IsCanceled = ?", 0).Where(
		"UserID", uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Count(&count)
	for i, order := range orders {
		r.db.Where("OrderID = ?", order.ID).Find(&orders[i].OrderedItems)
		for j, item := range orders[i].OrderedItems {
			r.db.Where("ID = ?", item.ProductID).First(&orders[i].OrderedItems[j].Product)
		}
	}
	return orders, int(count)
}
