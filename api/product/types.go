package product

import "pro05shopping/domain/product"

// 创建商品请求参数
type CreateProductRequest struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

// 创建商品响应参数
type CreateProductResponse struct {
	Message string `json:"message"`
}

// 删除商品请求参数
type DeleteProductRequest struct {
	SKU string `json:"sku"`
}

// 更新商品请求参数
type UpdateProductRequest struct {
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

// 类型转换 UpdateProductRequest to Product
func (p *UpdateProductRequest) ToProduct() *product.Product {
	return product.NewProduct(p.Name, p.Desc, p.Count, p.Price, p.CategoryID)
}
