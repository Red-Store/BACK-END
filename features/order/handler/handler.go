package handler

import (
	"MyEcommerce/features/order"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
}

func New(os order.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: os,
	}
}

func (handler *OrderHandler) CreateOrder(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newOrder := OrderRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	orderCore := RequestToCoreOrder(newOrder)
	items := []order.OrderItemCore{}
	payment, errInsert := handler.orderService.CreateOrder(userIdLogin, newOrder.CartIDs, orderCore, items)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert order", nil))
	}
	
	
	result := CoreToResponse(payment)

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", result))
}

func (handler *OrderHandler) GetOrderUser(c echo.Context) error  {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	results, errSelect := handler.orderService.GetOrderUser(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	userResult := make([]GetOrderUserResponse, len(results))
	for i, result := range results {
		userResult[i] = CoreToResponseOrderUser(result.Order, []order.OrderItemCore{result})
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", userResult))

}


func (handler *OrderHandler) GetOrderAdmin(c echo.Context) error  {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	results, errSelect := handler.orderService.GetOrderAdmin(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	response := CoreToResponseOrderAdmin(results)

	// Kirim response
	return c.JSON(http.StatusOK, response)
}