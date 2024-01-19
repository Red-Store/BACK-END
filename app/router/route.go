package router

import (
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/encrypts"
	"MyEcommerce/utils/middlewares"

	ud "MyEcommerce/features/user/data"
	uh "MyEcommerce/features/user/handler"
	us "MyEcommerce/features/user/service"

	pd "MyEcommerce/features/product/data"
	ph "MyEcommerce/features/product/handler"
	ps "MyEcommerce/features/product/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	productData := pd.New(db)
	productService := ps.New(productData)
	productHandlerAPI := ph.New(productService, cloudinaryUploader)

	// define routes/ endpoint ADMIN
  e.GET("admin/users", userHandlerAPI.GetAdminUserData, middlewares.JWTMiddleware())

	// define routes/ endpoint USERS
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
  e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	// define routes/ endpoint PRODUCTS
	e.POST("/products", productHandlerAPI.CreateProduct, middlewares.JWTMiddleware())
	e.GET("/products", productHandlerAPI.GetAllProduct)
	e.GET("/products/:id", productHandlerAPI.GetProductById)
	e.PUT("/products/:id", productHandlerAPI.UpdateProductById, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productHandlerAPI.DeleteProductById, middlewares.JWTMiddleware())
	e.GET("/users/products", productHandlerAPI.GetProductByUserId, middlewares.JWTMiddleware())
	e.GET("/products/search", productHandlerAPI.SearchProduct)
	// define routes/ endpoint CARTS


  // define routes/ endpoint ORDERS

}
