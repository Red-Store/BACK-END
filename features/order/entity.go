package order

import (
	cd "MyEcommerce/features/cart"
	ud "MyEcommerce/features/user"
	"time"
)

type OrderCore struct {
	// ID uint
	ID          string
	UserID      uint
	Address     string
	PaymentType string
	GrossAmount int
	Status      string
	VaNumber    int
	Bank        string
	User        ud.Core
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItemCore struct {
	OrderID   string
	CartID    uint
	Order     OrderCore
	Cart      cd.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderDataInterface interface {
	InsertOrder(userIdLogin int, inputOrder OrderCore, cartIds []uint) error
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	CreateOrder(userIdLogin int, cartIds []uint, inputOrder OrderCore) error
}
