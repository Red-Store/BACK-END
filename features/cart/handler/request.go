package handler

import (
	"MyEcommerce/features/cart"
)

type CartRequest struct {
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity" form:"quantity"`
}

func RequestToCore(input CartRequest, userIdLogin, productId uint) cart.Core {
	return cart.Core{
		UserID:    userIdLogin,
		Quantity:  input.Quantity,
		ProductID: productId,
	}
}
