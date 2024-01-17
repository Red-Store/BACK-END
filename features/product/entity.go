package product

import "time"

type Core struct {
	ID           uint
	Name         string
	Description  string
	Category     string
	Stock        int
	Price        int
	PhotoProduct string
	UserID       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductDataInterface interface {
	Insert(UserID int, input Core) error
	
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(input Core) error
}
