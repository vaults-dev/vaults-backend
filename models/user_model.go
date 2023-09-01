package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUIDBaseModel
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	return
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
