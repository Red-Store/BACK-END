package handler

import (
	"MyEcommerce/features/product"
	"time"
)

type ProductResponse struct {
	ID           uint      `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Description  string    `json:"description" form:"description"`
	Category     string    `json:"category" form:"category"`
	Stock        int       `json:"stock" form:"stock"`
	Price        int       `json:"price" form:"price"`
	PhotoProduct string    `json:"photo_product" form:"photo_product"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
	UserID       uint      `json:"user_id" form:"user_id"`
}

func CoreToResponse(data product.Core) ProductResponse {
	return ProductResponse{
		ID:           data.ID,
		Name:         data.Name,
		Description:  data.Description,
		Category:     data.Category,
		Stock:        data.Stock,
		Price:        data.Price,
		PhotoProduct: data.PhotoProduct,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		UserID:       data.UserID,
	}
}

func CoreToResponseList(data []product.Core) []ProductResponse {
	var results []ProductResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
