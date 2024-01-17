package data

import (
	"MyEcommerce/features/product"
	"MyEcommerce/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Description  string
	Category     string
	Stock        int
	Price        int
	PhotoProduct string
	UserID       int
	User         data.User
}

func CoreToModel(input product.Core) Product {
	return Product{
		Name:         input.Name,
		Description:  input.Description,
		Category:     input.Category,
		Stock:        input.Stock,
		Price:        input.Price,
		PhotoProduct: input.PhotoProduct,
		UserID:       input.UserID,
	}
}

func (p Product) ModelToCore() product.Core {
	return product.Core{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Category:     p.Category,
		Stock:        p.Stock,
		Price:        p.Price,
		PhotoProduct: p.PhotoProduct,
		UserID:       p.UserID,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}
