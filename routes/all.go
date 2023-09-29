package routes

import (
	AuthorController "gadgetify/controllers/author"
	ProductController "gadgetify/controllers/product"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {

	e.Use(middleware.Logger())              // logger in terminal
	e.Pre(middleware.RemoveTrailingSlash()) // remove trailing slash

	e.GET("/", AuthorController.GetAuthor)

	e.GET("/product", ProductController.GetAllProducts)
}
