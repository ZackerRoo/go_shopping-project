package product

import (
	pagination "pro05shopping/utils/pageination"
)

type Service struct {
	productRepository Repository
}

// 实例化
func NewService(productRepository Repository) *Service {
	productRepository.Migration()
	return &Service{
		productRepository: productRepository,
	}

}

// 获得所有商品分页
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page

}

// 创建商品
func (c *Service) CreateProduct(name string, desc string, count int, price float32, cid uint) error {
	newProduct := NewProduct(name, desc, count, price, cid)
	err := c.productRepository.Create(newProduct)
	return err
}

// 删除商品
func (c *Service) DeleteProduct(sku string) error {
	err := c.productRepository.Delete(sku)
	return err
}

// 更新商品
func (c *Service) UpdateProduct(product *Product) error {
	err := c.productRepository.Update(*product)
	return err
}

// 查找商品
func (c *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}
