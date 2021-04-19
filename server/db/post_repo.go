package db

import (
	"golang.org/x/crypto/bcrypt"
	"server/model"
	"server/utils"
	"server/utils/result"
)

type PostRepo struct {

}

func NewPostRepo() *PostRepo {
	return &PostRepo{}
}
func (ur *PostRepo) CheckExistViaEmail(email string) (code int, user *model.User) {
	var u model.User
	err := DB.Select("*").Where("email = ?", email).First(&u).Error
	if err != nil {
		return result.UserNotExist, nil
	}
	return result.UserExist, &u
}
func (ur *PostRepo) CheckExistViaID(id int) (code int, user *model.User) {
	var u model.User
	err := DB.Select("*").Where("id = ?", id).First(&u).Error
	if err != nil {
		return result.UserNotExist, nil
	}
	return result.UserExist, &u
}


func (ur *PostRepo) Insert(user model.User) (code int, message *string) {
	user.Password = utils.Encrypt(user.Password)
	err := DB.Create(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *PostRepo) DeleteVia(userID int) (code int, message *string) {
	var user model.User
	err := DB.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *PostRepo) Edit(user model.User) (code int, message *string) {
	err := DB.Model(&user).Where("id = ?", user.ID).Select("username", "avatar", "role").Updates(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *PostRepo) ChangePassword(user model.User) (code int, message *string) {
	user.Password = utils.Encrypt(user.Password)
	err := DB.Model(&user).Where("id = ?", user.ID).Select("password").Updates(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *PostRepo) GetVia(userID int) (code int, message *string, user model.APIUser) {
	var apiUser model.APIUser
	err := DB.Model(&model.User{}).Where("id = ?", userID).First(&apiUser).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, apiUser
	}
	return result.Success, nil, apiUser
}
func (ur *PostRepo)GetAllUsers(username string, pageSize int, pageNumber int) (int, *string, []model.APIUser, int64) {
	var users []model.APIUser
	var err error
	var total int64
	if username == "" {//select all users
		DB.Model(&[]model.User{}).Count(&total)
		if err = DB.Model(&[]model.User{}).Limit(pageSize).Offset(pageNumber).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	} else { // select all users with the same username
		DB.Model(&[]model.User{}).Where("username like ?", username).Count(&total)
		if err = DB.Model(&[]model.User{}).Where("username like ?", username).Limit(pageSize).Offset(pageNumber).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	}
	return result.Success, nil, users, total
}
func (ur *PostRepo) Login(email string, password string) (code int, message *string) {
	var user model.User
	if err := DB.Select("password").Where("email = ?", email).First(&user).Error; err != nil {
		return result.UserNotExist, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return result.UserPasswordNotRight, nil
	}
	return result.Success, nil
}
func (ur *PostRepo) DeleteAll() {
	DB.Exec("DELETE FROM users")
}