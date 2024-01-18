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

func (ps *productService) Create(userIdLogin int, input product.Core) error {
	err := ps.productData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}
	return nil
}
