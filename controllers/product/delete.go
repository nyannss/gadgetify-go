package product

import (
	"gadgetify/configs"
	BaseModel "gadgetify/models/base"
	ProductModel "gadgetify/models/product"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	var response BaseModel.BaseResponse

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		msg := "Invalid ID parameter"
		response.ErrorMessage = &msg
		return c.JSON(http.StatusBadRequest, response)
	}

	result := configs.DB.Delete(&ProductModel.Product{}, id)

	if result.RowsAffected == 0 {
		response.Status = http.StatusNotFound
		msg := "Data not Found"
		response.ErrorMessage = &msg
		response.Data = nil
		return c.JSON(response.Status, response)
	}

	msg := "Data successfully deleted"
	response.ErrorMessage = &msg
	response.Status = http.StatusAccepted

	return c.JSON(response.Status, response)
}
