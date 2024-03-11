package product

import (
	"errors"
)

var (
	ErrProductNotFound         = errors.New("商品没有找到")
	ErrProductStockIsNotEnough = errors.New("商品库存不足")
)
