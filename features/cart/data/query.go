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
	var cartGorm Cart

	tx := repo.db.Where("user_id = ? AND product_id = ?", userIdLogin, productId).First(&cartGorm)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			newCart := Cart{
				UserID:    uint(userIdLogin),
				ProductID: uint(productId),
				Quantity:  1,
			}
			return repo.db.Create(&newCart).Error
		}
		return tx.Error
	}
	cartGorm.Quantity++
	return repo.db.Save(&cartGorm).Error
}

// Select implements cart.CartDataInterface.
func (repo *cartQuery) Select(userIdLogin int) ([]cart.Core, error) {
	panic("unimplemented")
}

// Update implements cart.CartDataInterface.
func (repo *cartQuery) Update(userIdLogin int, cartId int, input cart.Core) error {
	panic("unimplemented")
}
