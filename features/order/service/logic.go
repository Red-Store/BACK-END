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
func (os *orderService) CreateOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore, items []order.OrderItemCore) (*order.OrderCore, error) {
	payment, err := os.orderData.InsertOrder(userIdLogin, cartIds, inputOrder, items)
	if err != nil {
		return nil, err
	}

	return payment, nil
}


// GetOrderUser implements order.OrderServiceInterface.
func (os *orderService) GetOrderUser(userIdLogin int) ([]order.OrderItemCore, error) {
	result, err := os.orderData.SelectOrderUser(userIdLogin)
	return result, err
}