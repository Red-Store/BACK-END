package product

import (
	"MyEcommerce/features/user/data"
	"time"
)

type Core struct {
	ID           uint
	Name         string
	Description  string
	Category     string
	Stock        int
	Price        int
	PhotoProduct string 
	UserID       uint
	User         data.User
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductDataInterface interface {
	Insert(userIdLogin int, input Core) error
	SelectAll() ([]Core, error)
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(userIdLogin int, input Core) error
	GettAll() ([]Core, error)
}
