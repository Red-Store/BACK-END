package data

import (
	"MyEcommerce/features/product"
	"errors"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

func (repo *productQuery) Insert(UserID int, input product.Core) error {
	productInputGorm := Product{
		UserID:       UserID,
		Name:         input.Name,
		Description:  input.Description,
		Category:     input.Category,
		Stock:        input.Stock,
		Price:        input.Price,
		PhotoProduct: input.PhotoProduct,
	}

	// simpan ke DB
	tx := repo.db.Create(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}
