package router

import (
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/encrypts"
	"MyEcommerce/utils/externalapi"
	"MyEcommerce/utils/middlewares"

	ud "MyEcommerce/features/user/data"
	uh "MyEcommerce/features/user/handler"
	us "MyEcommerce/features/user/service"

	pd "MyEcommerce/features/product/data"
	ph "MyEcommerce/features/product/handler"
	ps "MyEcommerce/features/product/service"

	cd "MyEcommerce/features/cart/data"
	ch "MyEcommerce/features/cart/handler"
	cs "MyEcommerce/features/cart/service"

	od "MyEcommerce/features/order/data"
	oh "MyEcommerce/features/order/handler"
	os "MyEcommerce/features/order/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	// midtransConfig := config.Midtrans{}
	// err := midtransConfig.LoadFromEnv("local.env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	midtrans := externalapi.New()
	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	productData := pd.New(db)
	productService := ps.New(productData)
	productHandlerAPI := ph.New(productService, cloudinaryUploader)

	cartData := cd.New(db)
	cartService := cs.New(cartData)
	cartHandlerAPI := ch.New(cartService)

	orderData := od.New(db)
	orderService := os.New(orderData)
	orderHandlerAPI := oh.New(orderService, midtrans)

	// define routes/ endpoint ADMIN
	e.GET("admin/users", userHandlerAPI.GetAdminUserData, middlewares.JWTMiddleware())

	// define routes/ endpoint USERS
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	// define routes/ endpoint PRODUCTS
	e.POST("/products", productHandlerAPI.CreateProduct, middlewares.JWTMiddleware())
	e.GET("/products", productHandlerAPI.GetAllProduct)
	e.GET("/products/:id", productHandlerAPI.GetProductById)
	e.PUT("/products/:id", productHandlerAPI.UpdateProductById, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productHandlerAPI.DeleteProductById, middlewares.JWTMiddleware())
	e.GET("/users/products", productHandlerAPI.GetProductByUserId, middlewares.JWTMiddleware())
	e.GET("/products/search", productHandlerAPI.SearchProduct)

	// define routes/ endpoint CARTS
	e.POST("/carts/:product_id", cartHandlerAPI.CreateCart, middlewares.JWTMiddleware())
	e.PUT("/carts/:cart_id", cartHandlerAPI.UpdateCart, middlewares.JWTMiddleware())
	e.DELETE("/carts/:cart_id", cartHandlerAPI.DeleteProductCart)
	e.GET("/carts", cartHandlerAPI.GetProductCart, middlewares.JWTMiddleware())

	// define routes/ endpoint ORDERS
	e.POST("/orders", orderHandlerAPI.CreateOrder, middlewares.JWTMiddleware())

}
