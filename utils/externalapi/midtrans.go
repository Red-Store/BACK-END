package externalapi

import (
	"MyEcommerce/app/config"
	"MyEcommerce/features/order"
	"errors"
	"fmt"
	"strconv"

	mid "github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransInterface interface {
	NewOrderPayment(data order.OrderCore, items []order.OrderItemCore) (*order.OrderCore, error)
}

type Midtrans struct {
	ApiKey string
	Env    mid.EnvironmentType
}

type midtrans struct {
	config Midtrans
	client coreapi.Client
}


func New() MidtransInterface {
	midConfig := Midtrans{
		ApiKey: config.MID_KEY,
	}

	// Parse MID_SANDBOX as an integer and determine the environment
	sandbox, err := strconv.Atoi(config.MID_SANDBOX)
	if err != nil {
		return &midtrans{}
	}
	if sandbox == 0 {
		midConfig.Env = mid.Production
	} else {
		midConfig.Env = mid.Sandbox
	}

	return &midtrans{
		config: midConfig,
		client: coreapi.Client{},
	}
}



// NewOrderPayment implements Midtrans.
func (pay *midtrans) NewOrderPayment(data order.OrderCore, items []order.OrderItemCore) (*order.OrderCore, error) {
	req := new(coreapi.ChargeReq)
	
	req.TransactionDetails = mid.TransactionDetails{
		OrderID:  data.ID,
		GrossAmt: int64(data.GrossAmount),
	}		


	var reqItem []mid.ItemDetails
	for _, item := range items {
		reqItem = append(reqItem, mid.ItemDetails{
			ID:    fmt.Sprintf("%d", item.CartID),
			Name:  item.Cart.Product.Name,
			Price: int64(item.Cart.Product.Price),
			Qty:   int32(item.Cart.Quantity),
		})
	}

	req.Items = &reqItem

	if data.PaymentType == "" {
		data.PaymentType = "bank_transfer"
	}

	switch data.Bank {
	case "bca":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBca,
		}
	case "bni":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBni,
		}
	case "bri":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBri,
		}

	default:
		return nil, errors.New("unsupported payment")

	}

	res, err := pay.client.ChargeTransaction(req)
	if err != nil {
		return nil, err
	}

	// Check the transaction status
	if res.StatusCode != "201" {
		return nil, errors.New(res.StatusMessage)
	}

	// Update the order data with the payment details
	data.VaNumber, _ = strconv.Atoi(res.VaNumbers[0].VANumber)
	data.PaymentType = res.PaymentType
	data.Status = res.TransactionStatus

	return &data, nil
}
