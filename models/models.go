package model

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Type    string
	Name    string
	Age     int
	Species string
	Gender  string
	Color   string
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
