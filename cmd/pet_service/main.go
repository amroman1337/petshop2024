package main

import (
	"github.com/amroman1337/petstore2024/database"
	"github.com/amroman1337/petstore2024/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title PetShop by Roman
// @version 1.0
// @description Backend practika

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.InitDB()
	routes.InitRoutes(e)
	routes.InitAuthRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
