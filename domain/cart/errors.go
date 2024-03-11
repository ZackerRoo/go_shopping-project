package cart

import (
	"errors"
)

var (
	ErrItemAlreadyExistInCart = errors.New("商品已经存在")
	ErrCountInvalid           = errors.New("数量不能是负值")
)
