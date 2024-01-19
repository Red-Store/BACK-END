package handler

import (
	"MyEcommerce/features/user"
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"
	"strconv"

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

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := RequestToCore(newUser)
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) GetUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	result, errSelect := handler.userService.GetById(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var userResult = CoreToResponse(result)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", userResult))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	fileData, err := c.FormFile("photo_profile")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
	}

	imageURL, err := handler.cld.UploadImage(fileData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image", nil))
	}

	userCore := UpdateRequestToCore(userData, imageURL)
	errUpdate := handler.userService.Update(userIdLogin, userCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
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

func (handler *UserHandler) GetAdminUserData(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	page, errPage := strconv.Atoi(c.QueryParam("page"))
	if errPage != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. page should be number", nil))
	}
	limit, errLimit := strconv.Atoi(c.QueryParam("limit"))
	if errLimit != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. limit should be number", nil))
	}

	result, errSelect := handler.userService.GetAdminUsers(userIdLogin, page, limit)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var userResult = CoreToResponseList(result)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", userResult))
}
