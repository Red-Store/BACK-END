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
	createdOrder, errInsert := handler.orderService.CreateOrder(newOrder.CartIDs, userIdLogin, &orderCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert order", nil))
	}
	

	items := []order.OrderItemCore{} 
	

	payment, err := handler.paymentMidtrans.NewOrderPayment(*createdOrder, items)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error create payment", err.Error()))
	}

	createdOrder.PaymentType = payment.PaymentType
	createdOrder.Status = payment.Status
	createdOrder.VaNumber = payment.VaNumber
	createdOrder.GrossAmount = payment.GrossAmount

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", payment))

}
