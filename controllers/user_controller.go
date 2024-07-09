package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/amroman1337/petstore2024/database"
	"github.com/amroman1337/petstore2024/models"
	"github.com/amroman1337/petstore2024/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @Summary SignUp
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User information"
// @Success 201 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
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
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
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
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
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

var jwtSecret = []byte("your-secret-key")

// @Summary Login a user
// @Description Login a user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param input body models.User true "User information"
// @Success 200 {object} map[string]string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/login [post]
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := services.AuthenticateUser(username, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": t})
}

// @Summary Logout a user
// @Description Logout a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/logout [post]
func Logout(c echo.Context) error {
	// Invalidate the token by setting it to an empty string or using a blacklist
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out"})
}

// @Summary Order a pet
// @Description Order a pet by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "Pet ID"
// @Success 200 {object} models.Pet
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /pets/order/{id} [post]
func OrderPet(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		var currentUser models.User
		if err := db.Where("username = ?", username).First(&currentUser).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, "User not found")
		}

		petID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid pet ID")
		}

		var pet models.Pet
		if err := db.First(&pet, petID).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Pet not found")
		}

		pet.OrderedBy = uint(currentUser.Userid)
		if err := db.Save(&pet).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to update pet order")
		}

		return c.JSON(http.StatusOK, pet)
	}
}
