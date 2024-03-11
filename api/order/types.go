package order

// 完成订单请求
type CompleteOrderRequest struct {
}

// 取消订单
type CancelOrderRequest struct {
	OrderId uint `json:"orderId"`
}
