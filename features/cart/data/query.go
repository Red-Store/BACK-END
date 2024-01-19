package data

import (
	"MyEcommerce/features/cart"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartDataInterface {
	return &cartQuery{
		db: db,
	}
}

// Delete implements cart.CartDataInterface.
func (repo *cartQuery) Delete(userIdLogin int, cartId int) error {
	panic("unimplemented")
}

// Insert implements cart.CartDataInterface.
func (repo *cartQuery) Insert(userIdLogin int, productId int) error {
	panic("unimplemented")
}

// Select implements cart.CartDataInterface.
func (repo *cartQuery) Select(userIdLogin int) ([]cart.Core, error) {
	panic("unimplemented")
}

// Update implements cart.CartDataInterface.
func (repo *cartQuery) Update(userIdLogin int, cartId int, input cart.Core) error {
	panic("unimplemented")
}
