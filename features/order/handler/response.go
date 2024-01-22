package handler

import "MyEcommerce/features/order"

type OrderResponse struct {
	OrderID string `json:"order_id" form:"order_id"`
	Address string `json:"address" form:"address"`
	Payment PaymentResponse
}

type PaymentResponse struct {
	StatusCode      string `json:"status_code" form:"status_code"`
	StatusMessage   string `json:"status_message" form:"status_message"`
	Status          string `json:"status" form:"status"`
	PaymentType     string `json:"payment_type" form:"payment_type"`
	GrossAmount     int    `json:"gross_amount" form:"gross_amount"`
	Bank            string `json:"bank" form:"bank"`
	VaNumber        int    `json:"va_number" form:"va_number"`
	TransactionId   string `json:"transaction_id" form:"transaction_id"`
	Currency        string `json:"currency" form:"currency"`
	TransactionTime string `json:"transaction_time" form:"transaction_time"`
	FraudStatus     string `json:"fraud_status" form:"fraud_status"`
}

func CoreToResponse(data *order.OrderCore) OrderResponse {
	var result = OrderResponse{
		OrderID: data.ID,
		Address: data.Address,
		Payment: PaymentResponse{
			StatusCode:      data.Payment.StatusCode,
			StatusMessage:   data.Payment.StatusMessage,
			Status:          data.Status,
			PaymentType:     data.PaymentType,
			GrossAmount:     data.GrossAmount,
			Bank:            data.Bank,
			VaNumber:        data.VaNumber,
			TransactionId:   data.Payment.TransactionId,
			Currency:        data.Payment.Currency,
			TransactionTime: data.Payment.TransactionTime,
			FraudStatus:     data.Payment.FraudStatus,
		},
	}
	return result
}
