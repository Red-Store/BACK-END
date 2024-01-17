package handler

import (
	"MyEcommerce/features/user"
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
	cld         cloudinary.CloudinaryUploaderInterface
}

func New(service user.UserServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *UserHandler {
	return &UserHandler{
		userService: service,
		cld:         cloudinaryUploader,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login "+err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}
