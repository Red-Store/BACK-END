package handler

import (
	"MyEcommerce/features/order"
	"MyEcommerce/utils/externalapi"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
	paymentMidtrans externalapi.MidtransInterface
}

func New(os order.OrderServiceInterface, mi externalapi.MidtransInterface) *OrderHandler {
	return &OrderHandler{
		orderService: os,
		paymentMidtrans: mi,
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
	errInsert := handler.orderService.CreateOrder(userIdLogin, newOrder.CartIDs, orderCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	items := []order.OrderItemCore{} 
	

	payment, err := handler.paymentMidtrans.NewOrderPayment(orderCore, items)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error create payment", err.Error()))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", payment))

}
