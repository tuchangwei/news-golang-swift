package model

import (
	"gorm.io/gorm"
)
//For these fields we use point, because point can be null.
//So we can distinguish if the client sends us includes the filed or not.
//If some fields are still nil after `ShouldBindJSON`, that means the client didn't send us these parameters.

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







