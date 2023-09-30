package main

import (
	"gadgetify/configs"
	"gadgetify/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// initialization environment
	configs.InitEnvironment()
	configs.InitDatabase()
	configs.InitCloudinary()

	// create new echo instance
	e := echo.New()

	// routes initialization
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
