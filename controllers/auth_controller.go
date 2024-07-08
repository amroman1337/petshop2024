package controllers

import (
	"net/http"
	"time"

	"github.com/amroman1337/petstore2024/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("your-secret-key")

// @Summary Login a user
// @Description Login a user with the input payload
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "User information"
// @Success 200 {object} map[string]string
// @Router /auth/login [post]
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
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func Logout(c echo.Context) error {
	// Invalidate the token by setting it to an empty string or using a blacklist
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out"})
}
