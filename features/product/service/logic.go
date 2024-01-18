package service

import (
	"MyEcommerce/features/product"
)

type productService struct {
	productData product.ProductDataInterface
}

func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

func (ps *productService) Create(userIdLogin int, input product.Core) error {
	err := ps.productData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}
	return nil
}

// GettAll implements product.ProductServiceInterface.
func (ps *productService) GetAll(page, limit int) ([]product.Core, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 8
	}

	products, err := ps.productData.SelectAll(page, limit)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetById implements product.ProductServiceInterface.
func (ps *productService) GetById(IdProduct int) (*product.Core, error) {
	result, err := ps.productData.SelectById(IdProduct)
	return result, err
}
