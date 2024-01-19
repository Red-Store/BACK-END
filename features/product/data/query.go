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

func (repo *productQuery) Insert(userIdLogin int, input product.Core) error {

	productInputGorm := CoreToModel(input)

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

// SelectAll implements product.ProductDataInterface.
func (repo *productQuery) SelectAll(page, limit int) ([]product.Core, error) {
	var products []Product
	err := repo.db.Limit(limit).Offset((page - 1) * limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	var productCores []product.Core
	for _, p := range products {
		productCores = append(productCores, p.ModelToCore())
	}

	return productCores, nil
}

// SelectById implements product.ProductDataInterface.
func (repo *productQuery) SelectById(IdProduct int) (*product.Core, error) {
	var productDataGorm Product
	tx := repo.db.Preload("User").Where("id = ?", IdProduct).First(&productDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := productDataGorm.ModelToCore()
	return &result, nil
}

// Update implements product.ProductDataInterface.
func (repo *productQuery) Update(userIdLogin int, input product.Core) error {
	product := Product{}
	tx := repo.db.Where("id = ? AND user_id = ?", input.ID, userIdLogin).First(&product)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errors.New("you do not have permission to edit this product")
		}
		return tx.Error
	}

	productInputGorm := CoreToModel(input)

	tx = repo.db.Model(&product).Updates(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Delete implements product.ProductDataInterface.
func (repo *productQuery) Delete(IdProduct int) error {
	tx := repo.db.Delete(&Product{}, IdProduct)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
