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

// Insert implements order.OrderDataInterface.
func (repo *orderQuery) InsertOrder(userIdLogin int,  inputOrder order.OrderCore, cartIds []uint) error {
	orderModel := CoreToModelOrder(inputOrder)
	orderModel.UserID = uint(userIdLogin)

	tx := repo.db.Create(&orderModel)
	if tx.Error != nil {
		return tx.Error
	}

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
