package data

import (
	"MyEcommerce/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (*userQuery) Insert(input user.Core) error {
	panic("unimplemented")
}

// SelectById implements user.UserDataInterface.
func (*userQuery) SelectById(userIdLogin int) (*user.Core, error) {
	panic("unimplemented")
}

// Update implements user.UserDataInterface.
func (*userQuery) Update(userIdLogin int, input user.Core) error {
	panic("unimplemented")
}

// Delete implements user.UserDataInterface.
func (*userQuery) Delete(userIdLogin int) error {
	panic("unimplemented")
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCore()
	return &result, nil
}
