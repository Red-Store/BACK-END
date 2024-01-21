package service

import (
	"MyEcommerce/features/order"
)

type orderService struct {
	orderData order.OrderDataInterface
}

func New(repo order.OrderDataInterface) order.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}

// CreateOrder implements order.OrderServiceInterface.
func (os *orderService) CreateOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore) error {
	err := os.orderData.InsertOrder(userIdLogin, inputOrder, cartIds)
	if err != nil {
		return err
	}

	return nil
}
