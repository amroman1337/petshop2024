package routes

import (
	"github.com/amroman1337/petstore2024/controllers"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/pets", controllers.CreatePet)
	e.GET("/pets/:id", controllers.GetPet)
	e.PUT("/pets/:id", controllers.UpdatePet)
	e.DELETE("/pets/:id", controllers.DeletePet)
	e.POST("/pets/filter", controllers.FilterPets(db))
	e.POST("/order/:id", controllers.OrderPet(db))
	users := e.Group("/users")
	users.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("your-secret-key"),
	}))
}

func InitAuthRoutes(e *echo.Echo) {
	e.POST("/users/login", controllers.Login)
	e.POST("/users/logout", controllers.Logout)
}
