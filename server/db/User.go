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
	Role *int `gorm:"not null;default:1" json:"role"`//1 normal, 2 admin
}
type APIUser struct {
	Username *string `json:"username"`
	Avatar *string `json:"avatar"`
	Role *int `gorm:"not null;default:1"`//1 normal, 2 admin
	ID   *uint `json:"id"`
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


func (u *User) Get() (code int, message *string, user APIUser) {
	var apiUser APIUser
	err := DB.Model(u).Where("id = ?", u.ID).First(&apiUser).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, apiUser
	}
	return result.Success, nil, apiUser
}
func GetAllUsers(username string, pageSize int, pageNumber int) (int, *string, []APIUser, int64) {
	var users []APIUser
	var err error
	if username == "" {//select all users
		if err = DB.Model(&[]User{}).Limit(pageSize).Offset(pageNumber-1).Find(&users).Error;
		err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	} else { // select all users with the same username
		if err = DB.Model(&[]User{}).Where("username like ?", username).Limit(pageSize).Offset(pageNumber).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	}
	var total int64
	DB.Model(&[]User{}).Count(&total)
	return result.Success, nil, users, total
}




