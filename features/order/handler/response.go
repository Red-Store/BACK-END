package handler

type OrderResponse struct {
	ID          uint   `json:"id" form:"id"`
	Address     string `json:"address" form:"address"`
	PaymentType string `json:"payment_type" form:"payment_type"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Status      string `json:"status" form:"status"`
	Bank        string `json:"bank" form:"bank"`
	VaNumber    int    `json:"va_number" form:"va_number"`
}
