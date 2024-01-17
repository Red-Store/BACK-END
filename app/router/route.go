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
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	// define routes/ endpoint
	e.POST("/login", userHandlerAPI.Login)
}
