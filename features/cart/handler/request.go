package handler

import (
	"MyEcommerce/features/cart"
)

type CartRequest struct {
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity" form:"quantity"`
}

type PutCartRequest struct {
	ID       uint
	UserID   uint
	Quantity int `json:"quantity" form:"quantity"`
}

func RequestToCore(input CartRequest, userIdLogin, productId uint) cart.Core {
	return cart.Core{
		UserID:    userIdLogin,
		Quantity:  input.Quantity,
		ProductID: productId,
	}
}

func RequestPutToCore(input PutCartRequest, userIdLogin, cartId uint) cart.Core {
	return cart.Core{
		UserID:   userIdLogin,
		Quantity: input.Quantity,
		ID:       cartId,
	}
}
