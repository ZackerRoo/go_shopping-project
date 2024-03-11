package category

import (
	"mime/multipart"
	"pro05shopping/utils/csv_helper"
	pagination "pro05shopping/utils/pageination"
)

type Service struct {
	r Repository
}

// 实例化商品分类service
func NewCategoryService(r Repository) *Service {
	// 生成表
	r.Migration()
	// 插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 创建分类
func (c *Service) Create(category *Category) error {
	existCity := c.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

// 批量创建分类
func (c *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := c.r.BulkCreate(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// 获得分页商品分类
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
