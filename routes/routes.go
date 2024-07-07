package routes

import (
	"github.com/amroman1337/petstore2024/controllers"
	"github.com/amroman1337/petstore2024/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/pets", controllers.CreatePet)
	e.GET("/pets/:id", controllers.GetPet)
	e.PUT("/pets/:id", controllers.UpdatePet)
	e.DELETE("/pets/:id", controllers.DeletePet)

	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("your-secret-key"),
	}))
	auth.POST("/order", controllers.OrderPet)
}
