package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type SignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
