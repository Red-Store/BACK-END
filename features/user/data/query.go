package data

import (
	"MyEcommerce/features/user"
	"errors"

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
func (repo *userQuery) Insert(input user.Core) error {
	dataGorm := CoreToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(userIdLogin int) (*user.Core, error) {
	var userDataGorm User
	tx := repo.db.First(&userDataGorm, userIdLogin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := userDataGorm.ModelToCore()
	return &result, nil
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
