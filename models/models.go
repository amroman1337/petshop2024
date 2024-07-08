package models

import (
	"gorm.io/gorm"
)

type Pet struct {
	Petid   int    `json:"Petid" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Species string `json:"species" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Color   string `json:"color" binding:"required"`
}

type User struct {
	Userid   int    `json:"Userid" binding:"required"`
	Username string `gorm:"unique" json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
	Email    string `gorm:"unique" json:"Email" binding:"required"`
}

type Order struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	PetID  uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
	Pet    Pet  `gorm:"foreignKey:PetID"`
}

type SearchCriteria struct {
	Type    string `json:"type"`
	Age     int    `json:"age"`
	Species string `json:"species"`
	Gender  string `json:"gender"`
	Color   string `json:"color"`
}
