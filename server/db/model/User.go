package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"not null" json:"email"`
	Username string `json:"username"`
	Password string `gorm:"not null" validate:"required,min=6,max=120" json:"password"`
}
