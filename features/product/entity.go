package product

import (
	"MyEcommerce/features/user"
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
	User         user.Core
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductDataInterface interface {
	Insert(userIdLogin int, input Core) error
	SelectAll(page, limit int) ([]Core, error)
	SelectById(IdProduct int) (*Core, error)
	Update(userIdLogin int, input Core) error
	Delete(IdProduct int) error
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(userIdLogin int, input Core) error
	GetAll(page, limit int) ([]Core, error)
	GetById(IdProduct int) (*Core, error)
	Update(userIdLogin int, input Core) error
	Delete(userIdLogin, IdProduct int) error
}
