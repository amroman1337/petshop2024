package main

import (
	"github.com/amroman1337/petstore2024/database"
	_ "github.com/amroman1337/petstore2024/docs"
	"github.com/amroman1337/petstore2024/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title PetShop by Roman
// @version 1.0
// @description Backend practika
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.InitDB()
	routes.InitRoutes(e, database.DB)
	routes.InitAuthRoutes(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
