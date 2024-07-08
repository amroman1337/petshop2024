package routes

import (
	"github.com/amroman1337/petstore2024/controllers"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/pets", controllers.CreatePet)
	e.GET("/pets/:id", controllers.GetPet)
	e.PUT("/pets/:id", controllers.UpdatePet)
	e.DELETE("/pets/:id", controllers.DeletePet)

	users := e.Group("/users")
	users.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("your-secret-key"),
	}))
}

func InitAuthRoutes(e *echo.Echo) {
	e.POST("/users/login", controllers.Login)
	e.POST("/users/logout", controllers.Logout)
}
