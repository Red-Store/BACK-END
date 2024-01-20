package service

import (
	"MyEcommerce/features/product"
	"errors"
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
	if input.Name == "" {
		return errors.New("nama produk tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga produk harus lebih besar dari 0")
	}

	err := ps.productData.Insert(userIdLogin, input)
	if err != nil {
		return err
	}
	
	return nil
}

// GettAll implements product.ProductServiceInterface.
func (ps *productService) GettAll() ([]product.Core, error) {
	products, err := ps.productData.SelectAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
