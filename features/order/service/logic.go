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
func (os *orderService) CreateOrder(cartIds []uint, userIdLogin int, inputOrder *order.OrderCore) (*order.OrderCore, error) {
	err := os.orderData.InsertOrder(inputOrder, userIdLogin, cartIds)
	if err != nil {
		return nil, err
	}

	return inputOrder, nil
}