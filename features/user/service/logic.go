package service

import (
	"MyEcommerce/features/user"
	"MyEcommerce/utils/encrypts"
	"MyEcommerce/utils/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	if input.Role == "" {
		input.Role = "user"
	}

	err := service.userData.Insert(input)
	return err
}

// GetById implements user.UserServiceInterface.
func (*userService) GetById(userIdLogin int) (*user.Core, error) {
	panic("unimplemented")
}

// Update implements user.UserServiceInterface.
func (*userService) Update(userIdLogin int, input user.Core) error {
	panic("unimplemented")
}

// Delete implements user.UserServiceInterface.
func (*userService) Delete(userIdLogin int) error {
	panic("unimplemented")
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil, "", err
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}
	log.Println("id user:", data.ID)
	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}
