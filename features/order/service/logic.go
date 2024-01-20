package service

import (
	"MyEcommerce/features/order"
)

type orderService struct {
	orderData order.OrderDataInterface
}

// Create implements order.OrderServiceInterface.
func (*orderService) Create(userIdLogin int, cartId int, input order.OrderCore) error {
	panic("unimplemented")
}

func New(repo order.OrderDataInterface) order.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}
