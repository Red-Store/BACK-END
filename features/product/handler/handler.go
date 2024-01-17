package handler

import (
	"MyEcommerce/features/product"
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
	cld            cloudinary.CloudinaryUploaderInterface
}

func New(ps product.ProductServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
		cld: cloudinaryUploader,
	}
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
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
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image", nil))
	}

	
	productCore := RequestToCore(newProduct, imageURL)
	productCore.PhotoProduct = imageURL

	errInsert := handler.productService.Create(productCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}
