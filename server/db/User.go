package db

import (
	"gorm.io/gorm"
	"server/utils/result"
)
//For these fields we use point, because point can be null.
//So we can distinguish if the client sends us includes the filed or not.
//If some fields are still nil after `ShouldBindJSON`, that means the client didn't send us these parameters.

type User struct {
	gorm.Model
	Email *string `gorm:"not null" validate:"required,email" json:"email"`
	Username *string `json:"username"`
	Password *string `gorm:"not null" validate:"required,min=6,max=120" json:"password"`
	Avatar *string `json:"avatar"`
	Role *int `gorm:"not null;default:1"`//1 normal, 2 admin
}


func (u *User) CheckExistViaEmail() (code int) {
	err := DB.Select("*").Where("email = ?", u.Email).First(u).Error
	if err != nil {
		return result.UserNotExist
	}
	return result.UserExist
}
func (u *User) CheckExistViaID() (code int) {
	err := DB.Select("*").Where("id = ?", u.ID).First(u).Error
	if err != nil {
		return result.UserNotExist
	}
	return result.UserExist
}

func (u *User) Insert() (code int, message *string) {
	err := DB.Create(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) Delete() (code int, message *string) {
	err := DB.Where("id = ?", u.ID).Delete(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) Edit() (code int, message *string) {
	err := DB.Model(u).Where("id = ?", u.ID).Select("username", "avatar", "role").Updates(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) ChangePassword() (code int, message *string) {
	err := DB.Model(u).Where("id = ?", u.ID).Select("password").Updates(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}


func (u *User) Get() (code int, message *string) {
	err := DB.Select("username", "avatar", "role").Where("id = ?", u.ID).First(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}




