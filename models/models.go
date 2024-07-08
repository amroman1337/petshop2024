package models

type Pet struct {
	Type    string `json:"type" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Species string `json:"species" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Color   string `json:"color" binding:"required"`
}

type User struct {
	Username string `gorm:"unique" json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
