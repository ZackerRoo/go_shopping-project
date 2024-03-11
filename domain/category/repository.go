package category

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

func (r *Repository) Create(category *Category) error {
	result := r.db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "电子产品", Desc: "手机、电脑、平板电脑等"},
		{Name: "服装", Desc: "男装、女装、童装等"},
	}

	for _, c := range categories {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

func (r *Repository) GetByName(name string) []Category {
	var category []Category
	r.db.Where("name = ?", name).First(&category)
	return category
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
}

func (r *Repository) BulkCreate(category []*Category) (int, error) {
	var count int64
	err := r.db.Create(&category).Count(&count).Error
	return int(count), err
}
