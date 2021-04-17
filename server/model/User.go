package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"not null" validate:"required,email" json:"email"`
	Username string `json:"username"`
	Password string `gorm:"not null" validate:"required,min=6,max=120" json:"password"`
	Avatar string `json:"avatar"`
	Role int `gorm:"not null;default:1" json:"role"`//1 normal, 2 admin
}
type APIUser struct {
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Role int `gorm:"not null;default:1" json:"role"`//1 normal, 2 admin
	ID   uint `json:"id"`
}







