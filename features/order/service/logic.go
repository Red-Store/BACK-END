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

// GetOrderAdmin implements order.OrderServiceInterface.
func (os *orderService) GetOrderAdmin(page, limit int) ([]order.OrderItemCore, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	result, err := os.orderData.SelectOrderAdmin(page, limit)
	return result, err
}

// CancleOrder implements order.OrderServiceInterface.
func (os *orderService) CancleOrder(userIdLogin int, orderId string, orderCore order.OrderCore) error {
	if orderCore.Status == "" {
		orderCore.Status = "cancelled"
	}

	err := os.orderData.CancleOrder(userIdLogin, orderId, orderCore)
	return err
}
