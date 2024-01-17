package handler

import (
	"MyEcommerce/features/product"
	"mime/multipart"
)

type ProductRequest struct {
	Name         string                `json:"name" form:"name"`
	Description  string                `json:"description" form:"description"`
	Category     string                `json:"category" form:"category"`
	Stock        int                   `json:"stock" form:"stock"`
	Price        int                   `json:"price" form:"price"`
	PhotoProduct *multipart.FileHeader `json:"photo_product" form:"photo_product"`
}

func RequestToCore(input ProductRequest, imageURL string) product.Core {
	return product.Core{
		Name:         input.Name,
		Description:  input.Description,
		Category:     input.Category,
		Stock:        input.Stock,
		Price:        input.Price,
		PhotoProduct: imageURL,
	}
}
