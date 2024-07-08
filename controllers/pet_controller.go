package controllers

import (
	"net/http"
	"strconv"

	"github.com/amroman1337/petstore2024/database"
	"github.com/amroman1337/petstore2024/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @Summary Create new pet
// @Description Create a new pet with the input payload
// @Tags pets
// @Accept  json
// @Produce  json
// @Param input body models.Pet true "Pet information"
// @Success 201 {object} models.Pet
// @Router /pets [post]
func CreatePet(c echo.Context) error {
	pet := new(models.Pet)
	if err := c.Bind(pet); err != nil {
		return err
	}
	database.DB.Create(&pet)
	return c.JSON(http.StatusCreated, pet)
}

// @Summary Get pet by ID
// @Description Get pet information by ID
// @Tags pets
// @Accept  json
// @Produce  json
// @Param id path int true "Pet ID"
// @Success 200 {object} models.Pet
// @Router /pets/{id} [get]
func GetPet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pet models.Pet
	database.DB.First(&pet, id)
	return c.JSON(http.StatusOK, pet)
}

// @Summary Update pet by ID
// @Description Update pet information by ID
// @Tags pets
// @Accept  json
// @Produce  json
// @Param id path int true "Pet ID"
// @Param pet body models.Pet true "Pet information"
// @Success 200 {object} models.Pet
// @Router /pets/{id} [put]
func UpdatePet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pet := new(models.Pet)
	if err := c.Bind(pet); err != nil {
		return err
	}
	database.DB.Model(&models.Pet{}).Where("id = ?", id).Updates(pet)
	return c.JSON(http.StatusOK, pet)
}

// @Summary Delete pet by ID
// @Description Delete pet by ID
// @Tags pets
// @Accept  json
// @Produce  json
// @Param id path int true "Pet ID"
// @Success 204 "No Content"
// @Router /pets/{id} [delete]
func DeletePet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Pet{}, id)
	return c.NoContent(http.StatusNoContent)
}

// @Summary Filter pets
// @Description Filter pets by type, age, species, gender, and color
// @Tags pets
// @Accept  json
// @Produce  json
// @Param type query string false "Type of pet"
// @Param age query int false "Age of pet"
// @Param species query string false "Species of pet"
// @Param gender query string false "Gender of pet"
// @Param color query string false "Color of pet"
// @Success 200 {array} models.Pet
// @Router /pets/filter [get]
func FilterPets(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var pets []models.Pet
		query := db

		if c.QueryParam("type") != "" {
			query = query.Where("type = ?", c.QueryParam("type"))
		}
		if c.QueryParam("age") != "" {
			query = query.Where("age = ?", c.QueryParam("age"))
		}
		if c.QueryParam("species") != "" {
			query = query.Where("species = ?", c.QueryParam("species"))
		}
		if c.QueryParam("gender") != "" {
			query = query.Where("gender = ?", c.QueryParam("gender"))
		}
		if c.QueryParam("color") != "" {
			query = query.Where("color = ?", c.QueryParam("color"))
		}

		query.Find(&pets)
		return c.JSON(http.StatusOK, pets)
	}
}
