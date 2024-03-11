package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	p.SKU = uuid.New().String()
	return
}
