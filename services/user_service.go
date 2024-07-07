package services

import (
	"github.com/amroman1337/petstore2024/database"
	"github.com/amroman1337/petstore2024/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}
	result := database.DB.Create(&user)
	return result.Error
}

func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
