package service

import (
	"MyEcommerce/features/cart"
)

type cartService struct {
	cartData cart.CartDataInterface
}

// dependency injection
func New(repo cart.CartDataInterface) cart.CartServiceInterface {
	return &cartService{
		cartData: repo,
	}
}

// Create implements cart.CartServiceInterface.
func (cs *cartService) Create(userIdLogin int, productId int) error {
	err := cs.cartData.Insert(userIdLogin, productId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCart implements cart.CartServiceInterface.
func (cs *cartService) DeleteCart(userIdLogin int, cartId int) error {
	panic("unimplemented")
}

// Get implements cart.CartServiceInterface.
func (cs *cartService) Get(userIdLogin int) ([]cart.Core, error) {
	panic("unimplemented")
}

// UpdateCart implements cart.CartServiceInterface.
func (cs *cartService) UpdateCart(userIdLogin int, cartId int, input cart.Core) error {
	panic("unimplemented")
}
