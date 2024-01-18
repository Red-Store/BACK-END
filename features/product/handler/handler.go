package handler

import (
	"MyEcommerce/features/product"
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
	cld            cloudinary.CloudinaryUploaderInterface
}

func New(ps product.ProductServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
		cld:            cloudinaryUploader,
	}
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	fileHeader, err := c.FormFile("photo_product")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
	}

	imageURL, err := handler.cld.UploadImage(fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image, jenis file salah", nil))
	}

	productCore := RequestToCore(newProduct, imageURL, uint(userIdLogin))
	productCore.PhotoProduct = imageURL

	errInsert := handler.productService.Create(userIdLogin, productCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *ProductHandler) GetAllProduct(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	products, startIndex, endIndex, err := handler.productService.GettAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error get data", nil))
	}

	response := map[string]interface{}{
		"products":   products,
		"startIndex": startIndex,
		"endIndex":   endIndex,
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success get data", response))
}
