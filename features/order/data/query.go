package data

import (
	"MyEcommerce/features/order"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderDataInterface {
	return &orderQuery{
		db: db,
	}
}

// InsertOrder implements order.OrderDataInterface.
func (repo *orderQuery) InsertOrder(inputOrder *order.OrderCore, userIdLogin int, cartIds []uint) error {
	orderModel := CoreToModelOrder(*inputOrder)
	orderModel.UserID = uint(userIdLogin)

	orderModel.PaymentType = inputOrder.PaymentType
	orderModel.Status = inputOrder.Status
	orderModel.VaNumber = inputOrder.VaNumber

	tx := repo.db.Create(&orderModel)
	if tx.Error != nil {
		return tx.Error
	}

	inputOrder.ID = orderModel.ID

	for _, cartId := range cartIds {
		orderItem := OrderItem{
			OrderID: orderModel.ID,
			CartID:  cartId,
		}

		if err := repo.db.Create(&orderItem).Error; err != nil {
			return err
		}
	}

	return nil
}