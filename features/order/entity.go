package order

import (
	cd "MyEcommerce/features/cart"
	ud "MyEcommerce/features/user"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OrderCore struct {
	ID          uuid.UUID 
	UserID      uint      
	Address     string
	PaymentType string
	GrossAmount int
	Status      string
	VaNumber    int
	Bank string
	User        ud.Core
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItemCore struct {
	gorm.Model
	OrderID   uint
	CartID    uint
	Order     OrderCore
	Cart      cd.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderDataInterface interface {
	Insert(userIdLogin, cartId int, input OrderCore) error
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	Create(userIdLogin, cartId int, input OrderCore) error
}
