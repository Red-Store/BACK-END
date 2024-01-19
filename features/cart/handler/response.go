package handler

import (
	"MyEcommerce/features/cart"
	"MyEcommerce/features/product/handler"
)

type CartResponse struct {
	ID       uint `json:"id" form:"id"`
	Quantity int  `json:"quantity" form:"quantity"`
	Products handler.CartProductResponse
}

func CoreToResponse(data cart.Core) CartResponse {	
	productResponse := handler.CartProductResponse{
		Name:         data.Product.Name,
		Price:        data.Product.Price,
		PhotoProduct: data.Product.PhotoProduct,
	}

	return CartResponse{
		ID:       data.ID,
		Quantity: data.Quantity,
		Products: productResponse,
	}
}

func CoreToResponseList(data []cart.Core) []CartResponse {
	var results []CartResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
