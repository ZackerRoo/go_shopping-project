package order

import "errors"

var (
	ErrEmptyCartFound       = errors.New("购物车是空的")
	ErrInvalidOrderID       = errors.New("无效的订单ID")
	ErrCancelDurationPassed = errors.New("无法取消订单，订单取消时间已过")
	ErrNotEnoughStock       = errors.New("库存不足")
)
