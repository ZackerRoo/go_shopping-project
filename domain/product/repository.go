package product

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		log.Print(err)
	}
}

// 更新
func (r *Repository) Update(updateProduct Product) error {
	savedProduct, err := r.FindBySKU(updateProduct.SKU)
	if err != nil {
		return err
	}
	err = r.db.Model(&savedProduct).Updates(updateProduct).Error
	return err
}

// 搜索返回分页结果
func (r *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where("IsDeleted = ?", false).Where(
		"Name LIKE ? OR SKU Like ?", convertedStr,
		convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

// 根据sku查找
func (r *Repository) FindBySKU(sku string) (*Product, error) {
	var product *Product
	err := r.db.Where("IsDeleted = ?", 0).Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

// 创建
func (r *Repository) Create(p *Product) error {
	result := r.db.Create(p)

	return result.Error
}

// 查询所有商品
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	r.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

// 根据sku删除
func (r *Repository) Delete(sku string) error {
	currentProduct, err := r.FindBySKU(sku)
	if err != nil {
		return err
	}
	currentProduct.IsDeleted = true

	err = r.db.Save(currentProduct).Error
	return err
}
