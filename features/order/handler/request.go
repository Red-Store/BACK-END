package handler

import (
	"MyEcommerce/features/order"

	"github.com/google/uuid"
)

type OrderRequest struct {
	ID          string
	Address     string `json:"address" form:"address"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Bank        string `json:"bank" form:"bank"`
	CartIDs     []uint `json:"cart_ids" form:"cart_ids"`
}

type CancleOrderRequest struct {
	Status string `json:"status"`
}

type WebhoocksRequest struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
}

// type WebhoocksRequest struct {
// 	TransactionID     string `json:"transaction_id"`
// 	StatusCode        string `json:"status_code"`
// 	StatusMessage     string `json:"status_message"`
// 	TransactionStatus string `json:"transaction_status"`
// 	PaymentType       string `json:"payment_type"`
// 	GrossAmount       string `json:"gross_amount"`
// 	Bank              string `json:"bank"`
// 	VaNumber          string `json:"va_number"`
// 	Currency          string `json:"currency"`
// 	TransactionTime   string `json:"transaction_time"`
// 	FraudStatus       string `json:"fraud_status"`
// }

func RequestToCoreOrder(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		ID:          uuid.New().String(),
		Address:     input.Address,
		GrossAmount: input.GrossAmount,
		Bank:        input.Bank,
	}
}

// func RequestToCoreWebhocks(input WebhoocksRequest) order.Webhoocks {
// 	return order.Webhoocks{
// 		TransactionID:     input.TransactionID,
// 		StatusCode:        input.StatusCode,
// 		StatusMessage:     input.StatusMessage,
// 		TransactionStatus: input.TransactionStatus,
// 		PaymentType:       input.PaymentType,
// 		GrossAmount:       input.GrossAmount,
// 		Bank:              input.Bank,
// 		VaNumber:          input.VaNumber,
// 		Currency:          input.Currency,
// 		TransactionTime:   input.TransactionTime,
// 		FraudStatus:       input.FraudStatus,
// 	}
// }

func CancleRequestToCoreOrder(input CancleOrderRequest) order.OrderCore {
	return order.OrderCore{
		Status: input.Status,
	}
}

func WebhoocksRequestToCore(input WebhoocksRequest) order.OrderCore {
	return order.OrderCore{
		ID:     input.OrderID,
		Status: input.TransactionStatus,
	}
}
