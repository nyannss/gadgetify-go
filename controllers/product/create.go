package product

import (
	"context"
	"database/sql"
	"gadgetify/configs"
	"gadgetify/models/base"
	CategoryModel "gadgetify/models/category"
	ProductModel "gadgetify/models/product"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func CreateNew(c echo.Context) error {
	var response base.BaseResponse
	product := new(ProductModel.Product)

	file, err := c.FormFile("image")
	if err == nil {
		if file.Size > 2*1024*1024 { // 2mb
			response.Status = http.StatusBadRequest
			msg := "File size exceeds limit"
			response.ErrorMessage = &msg
			return c.JSON(response.Status, response)
		}

		allowedFormats := []string{"image/jpeg", "image/png"}
		validFormat := false
		for _, format := range allowedFormats {
			if format == file.Header.Get("Content-Type") {
				validFormat = true
				break
			}
		}

		if !validFormat {
			response.Status = http.StatusBadRequest
			msg := "Invalid file format"
			response.ErrorMessage = &msg

			return echo.NewHTTPError(http.StatusBadRequest, response)
		}

		src, err := file.Open()
		if err != nil {
			response.Status = http.StatusInternalServerError
			msg := err.Error()
			response.ErrorMessage = &msg

			return echo.NewHTTPError(response.Status, response)
		}
		defer src.Close()

		uploadResult, err := configs.Cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: "gadgetify"})

		if err != nil {
			response.Status = http.StatusInternalServerError
			msg := err.Error()
			response.ErrorMessage = &msg

			return echo.NewHTTPError(response.Status, response)
		}

		product.ImageURL = &uploadResult.SecureURL
	}

	if err := c.Bind(product); err != nil {
		response.Status = http.StatusBadRequest
		errMsg := err.Error()
		response.ErrorMessage = &errMsg
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		response.Status = http.StatusBadRequest
		errMsg := err.Error()
		response.ErrorMessage = &errMsg
		return c.JSON(http.StatusBadRequest, response)
	}

	categoryID, err := strconv.Atoi(c.FormValue("category_id"))
	if err == nil {
		response.Status = http.StatusBadRequest
		var category CategoryModel.Category
		errCat := configs.DB.First(&category, categoryID)
		if errCat.Error != nil {
			errMsg := "Failed to retrieve category"
			response.ErrorMessage = &errMsg
			return c.JSON(http.StatusBadRequest, response)
		}
		product.CategoryID = sql.NullString{String: category.ID, Valid: true}
	}

	if err := configs.DB.Create(product).Error; err != nil {
		response.Status = http.StatusInternalServerError
		errMsg := "Failed to create product"
		response.ErrorMessage = &errMsg
		return c.JSON(response.Status, response)
	}

	// Retrieve the related category
	category := &CategoryModel.Category{}
	if err := configs.DB.First(&category, product.CategoryID).Error; err != nil {
		errMsg := "Failed to retrieve related category"
		response.ErrorMessage = &errMsg
		response.Status = http.StatusInternalServerError
		return c.JSON(response.Status, response)
	}

	// Attach the related category to the product
	product.Category = *category

	response.Status = http.StatusCreated
	response.Data = product
	return c.JSON(response.Status, response)
}
