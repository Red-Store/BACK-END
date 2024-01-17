package service

import "MyEcommerce/features/product"

type productService struct {
	productData product.ProductDataInterface
}

func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

func (ps *productService) Create(input product.Core) error {
	err := ps.productData.Insert(input.UserID, input)
	if err != nil {
		return err
	}
	return nil
}
