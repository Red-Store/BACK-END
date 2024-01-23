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
	Payment     Payment
	OrderItems []OrderItemCore
}

type OrderItemCore struct {
	OrderID   string
	CartID    uint
	Order     OrderCore
	Cart      cd.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	StatusCode      string
	StatusMessage   string
	TransactionId   string
	Currency        string
	TransactionTime string
	FraudStatus     string
}

// interface untuk Data Layer
type OrderDataInterface interface {
	InsertOrder(userIdLogin int, cartIds []uint, inputOrder OrderCore, items []OrderItemCore) (*OrderCore, error)
	SelectOrderUser(userIdLogin int) ([]OrderItemCore, error)
	SelectOrderAdmin(userIdLogin int) ([]OrderItemCore, error)
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	CreateOrder(userIdLogin int, cartIds []uint, inputOrder OrderCore, items []OrderItemCore) (*OrderCore, error)
	GetOrderUser(userIdLogin int) ([]OrderItemCore, error)
	GetOrderAdmin(userIdLogin int) ([]OrderItemCore, error)
}
