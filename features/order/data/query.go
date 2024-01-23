package data

import (
	"MyEcommerce/features/order"
	"MyEcommerce/utils/externalapi"

	"gorm.io/gorm"
)

type orderQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func New(db *gorm.DB, mi externalapi.MidtransInterface) order.OrderDataInterface {
	return &orderQuery{
		db:              db,
		paymentMidtrans: mi,
	}
}

// InsertOrder implements order.OrderDataInterface.
func (repo *orderQuery) InsertOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore, items []order.OrderItemCore) (*order.OrderCore, error) {
	payment, errPay := repo.paymentMidtrans.NewOrderPayment(inputOrder, items)
	if errPay != nil {
		return nil, errPay
	}

	orderModel := CoreToModelOrder(inputOrder)
	orderModel.UserID = uint(userIdLogin)

	tx := repo.db.Create(&orderModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	inputOrder.ID = orderModel.ID

	for _, cartId := range cartIds {
		orderItem := OrderItem{
			OrderID: orderModel.ID,
			CartID:  cartId,
		}

		if err := repo.db.Create(&orderItem).Error; err != nil {
			return nil, err
		}
	}

	return payment, nil
}
