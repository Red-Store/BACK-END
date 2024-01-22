package handler

import (
	"MyEcommerce/features/order"

	"github.com/google/uuid"
)

type OrderRequest struct {
	ID          string
	Address     string `json:"address" form:"address"`
	PaymentType string `json:"payment_type" form:"payment_type"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Status      string `json:"status" form:"status"`
	Bank        string `json:"bank" form:"bank"`
	VaNumber    int    `json:"va_number" form:"va_number"`
	CartIDs     []uint `json:"cart_ids" form:"cart_ids"`
}

func RequestToCoreOrder(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		ID:          uuid.New().String(),
		Address:     input.Address,
		PaymentType: input.PaymentType,
		GrossAmount: input.GrossAmount,
		Status:      input.Status,
		Bank:        input.Bank,
		VaNumber:    input.VaNumber,
	}
}
