package controllers

import (
	"net/http"
	"strconv"

	"github.com/amroman1337/petstore2024/database"
	"github.com/amroman1337/petstore2024/models"

	"github.com/labstack/echo/v4"
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
