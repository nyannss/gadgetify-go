package product

import (
	"gadgetify/configs"
	BaseModel "gadgetify/models/base"
	ProductModel "gadgetify/models/product"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	var (
		products  []ProductModel.Product
		response  BaseModel.BaseResponse
		count     int64
		perPage   int64
		totalPage float64
		page      int
	)

	// Get total count of products
	configs.DB.Model(&ProductModel.Product{}).Count(&count)

	// Handle 'per_page' parameter
	perPage, err := strconv.ParseInt(c.QueryParam("per_page"), 10, 64)
	if err != nil || perPage < 1 {
		perPage = 10
	} else if perPage > 50 {
		perPage = 50
	}

	// Calculate total pages
	totalPage = math.Ceil(float64(count) / float64(perPage))

	// Handle 'page' parameter
	page, _ = strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	} else if page > int(totalPage) {
		page = int(totalPage)
	}

	// Fetch products with pagination and category preloading
	result := configs.DB.Limit(int(perPage)).Preload("Category").Find(&products)

	if result.Error != nil {
		response.Status = http.StatusInternalServerError
		errorMsg := "Internal Server Error"
		response.ErrorMessage = &errorMsg
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = http.StatusOK

	// Prepare pagination metadata
	pagination := BaseModel.MetaPagination{
		CurrentPage: uint16(page),
		PerPage:     uint(perPage),
		TotalData:   uint32(count),
		TotalPage:   uint16(totalPage),
	}

	// Populate response data
	response.Data = BaseModel.BasePagination{
		Meta:  pagination,
		Items: products,
	}

	return c.JSON(http.StatusOK, response)
}
