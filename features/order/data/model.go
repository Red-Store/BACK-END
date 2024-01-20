package data

import (
	cd "MyEcommerce/features/cart/data"
	"MyEcommerce/features/order"
	ud "MyEcommerce/features/user/data"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID      uint      `gorm:"foreignKey:UserID"`
	User        ud.User
	Address     string
	PaymentType string
	GrossAmount int
	Status      string
	Bank        string
	VaNumber    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type OrderItem struct {
	gorm.Model
	OrderID uint
	Order   Order
	CartID  uint
	Cart    cd.Cart
}

func CoreToModelOrder(input order.OrderCore) Order {
	return Order{
		UserID:      input.UserID,
		Address:     input.Address,
		PaymentType: input.PaymentType,
		GrossAmount: input.GrossAmount,
		Status:      input.Status,
		VaNumber:    input.VaNumber,
		Bank:        input.Bank,
	}
}

func (o Order) ModelToCoreOrder() order.OrderCore {
	return order.OrderCore{
		UserID:      o.UserID,
		ID:          o.ID,
		Address:     o.Address,
		PaymentType: o.PaymentType,
		GrossAmount: o.GrossAmount,
		Status:      o.Status,
		VaNumber:    o.VaNumber,
		Bank:        o.Bank,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}

func CoreToModelOrderItem(input order.OrderItemCore) OrderItem {
	return OrderItem{
		OrderID: input.OrderID,
		CartID:  input.CartID,
	}
}

func (ot OrderItem) ModelToCoreOrderItem() order.OrderItemCore {
	return order.OrderItemCore{
		OrderID:   ot.OrderID,
		CartID:    ot.CartID,
		CreatedAt: ot.CreatedAt,
		UpdatedAt: ot.UpdatedAt,
		Cart:      ot.Cart.ModelToCore(),
		Order:     ot.Order.ModelToCoreOrder(),
	}
}
