package controllers

import (
	"net/http"
	"strconv"

	"github.com/amroman1337/petstore2024/database"
	"github.com/amroman1337/petstore2024/models"
	"github.com/amroman1337/petstore2024/services"

	"github.com/labstack/echo/v4"
)

// @Summary SignUp
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User information"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	err := services.CreateUser(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

// @Summary Get user by ID
// @Description Get user information by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	database.DB.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

// @Summary Update user by ID
// @Description Update user information by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User information"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	database.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)
	return c.JSON(http.StatusOK, user)
}
