package author

import (
	"gadgetify/models/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Contributors struct {
	Alias  string `json:"alias"`
	Github string `json:"github"`
}

type DataResponse struct {
	Contributors interface{} `json:"contributors"`
}

func GetAuthor(c echo.Context) error {
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:       http.StatusOK,
		ErrorMessage: nil,
		Data: DataResponse{
			Contributors: []Contributors{{Alias: "nyannss", Github: "https://github.com/nyannss"}},
		},
	})
}
