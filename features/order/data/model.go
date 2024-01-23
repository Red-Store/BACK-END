package data

import (
	cd "MyEcommerce/features/cart/data"
	"MyEcommerce/features/order"

	ud "MyEcommerce/features/user/data"

	"gorm.io/gorm"
)

type Order struct {
	ID string `gorm:"type:varchar(36);primary_key" json:"id"`
	gorm.Model
	UserID      uint 
	User        ud.User
	Address     string
	PaymentType string
	GrossAmount int
	Status      string
	Bank        string
	VaNumber    int
}

type OrderItem struct {
	gorm.Model
	OrderID string
	Order   Order
	CartID  uint
	Cart    cd.Cart
}

func CoreToModelOrder(input order.OrderCore) Order {
	return Order{
		ID:          input.ID,
		UserID:      input.UserID,
		Address:     input.Address,
		GrossAmount: input.GrossAmount,
		Status:      input.Status,
		VaNumber:    input.VaNumber,
		Bank:        input.Bank,
	}
}

func (o Order) ModelToCoreOrderUser() order.OrderCore {
	return order.OrderCore{
		ID:          o.ID,
		UserID:      o.UserID,
		Address:     o.Address,
		PaymentType: o.PaymentType,
		GrossAmount: o.GrossAmount,
		Status:      o.Status,
		VaNumber:    o.VaNumber,
		Bank:        o.Bank,
		User:        o.User.ModelToCore(),
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}

func (ot OrderItem) ModelToCoreOrderItemUser() order.OrderItemCore {
	return order.OrderItemCore{
		OrderID:   ot.OrderID,
		CartID:    ot.CartID,
		CreatedAt: ot.CreatedAt,
		UpdatedAt: ot.UpdatedAt,
		Cart:      ot.Cart.ModelToCore(),
		Order:     ot.Order.ModelToCoreOrderUser(),
	}
}