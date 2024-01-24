package service

import (
	"MyEcommerce/features/order"
	"errors"
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
	if len(cartIds) == 0 {
		return nil, errors.New("masukan barang anda")
	}
	if inputOrder.Address == "" {
		return nil, errors.New("masukan alamat anda")
	}

	if inputOrder.GrossAmount == 0 {
		return nil, errors.New("total pembayaran anda salah")
	}

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
func (os *orderService) GetOrderAdmin(userIdLogin int) ([]order.OrderItemCore, error) {
	result, err := os.orderData.SelectOrderAdmin(userIdLogin)
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
