package data

import (
	"MyEcommerce/features/order"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

// Insert implements order.OrderDataInterface.
func (*orderQuery) Insert(userIdLogin int, cartId int, input order.OrderCore) error {
	panic("unimplemented")
}

func New(db *gorm.DB) order.OrderDataInterface {
	return &orderQuery{
		db: db,
	}
}
