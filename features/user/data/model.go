package data

import (
	"MyEcommerce/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	Name         string
	UserName     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	Role         string
	PhotoProfile string
}

func CoreToModel(input user.Core) User {
	return User{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Password:     input.Password,
		Role:         input.Role,
		PhotoProfile: input.PhotoProfile,
	}
}

func (u User) ModelToCore() user.Core {
	return user.Core{
		ID:           u.ID,
		Name:         u.Name,
		UserName:     u.UserName,
		Email:        u.Email,
		Password:     u.Password,
		Role:         u.Role,
		PhotoProfile: u.PhotoProfile,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
