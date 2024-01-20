package handler

import (
	"MyEcommerce/features/order"
)

type OrderRequest struct {
	Address     string `json:"address" form:"address"`
	PaymentType string `json:"payment_type" form:"payment_type"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Status      string `json:"status" form:"status"`
	Bank        string `json:"bank" form:"bank"`
	VaNumber    int    `json:"va_number" form:"va_number"`
}

type OrderItemRequest struct {
	CartID  uint `json:"cart_id" form:"cart_id"`
	OrderID uint `json:"order_id" form:"order_id"`
}

func RequestToCoreOrder(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		Address:     input.Address,
		PaymentType: input.PaymentType,
		GrossAmount: input.GrossAmount,
		Status:      input.Status,
		Bank:        input.Bank,
		VaNumber:    input.VaNumber,
	}
}

func RequestToCoreOrderItem(input OrderItemRequest) order.OrderItemCore {
	return order.OrderItemCore{
		CartID:  input.CartID,
		OrderID: input.OrderID,
	}
}
