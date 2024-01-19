package handler

import "MyEcommerce/features/cart"

type CartHandler struct {
	cartService cart.CartServiceInterface
}

func New(cs cart.CartServiceInterface) *CartHandler {
	return &CartHandler{
		cartService: cs,
	}
}
