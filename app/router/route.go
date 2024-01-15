package router

import (
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/encrypts"

	ud "MyEcommerce/features/user/data"
	uh "MyEcommerce/features/user/handler"
	us "MyEcommerce/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := ud.New(db)
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	e.POST("/users", userHandlerAPI.CreateUser)
}
