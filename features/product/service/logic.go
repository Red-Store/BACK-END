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
func (ps *productService) GettAll(page, limit int) ([]product.Core, int, int, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 8
	}
	
	products, err := ps.productData.SelectAll(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	startIndex := (page - 1) * limit + 1
	endIndex := startIndex + len(products) - 1

	return products, startIndex, endIndex, nil
}
