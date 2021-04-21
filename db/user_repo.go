package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"server/utils"
	"server/utils/result"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}
func (ur *UserRepo) CheckExistViaEmail(email string) (code int, user User) {
	u := User{}
	u.Email = email
	err := DB.Select("*").First(&u).Error
	if err != nil {
		return result.UserNotExist, u
	}
	return result.UserExist, u
}
func (ur *UserRepo) CheckExistViaID(id int) (code int, user User) {
	var u = User{}
	u.ID = uint(id)
	err := DB.Select("*").First(&u).Error
	if err != nil {
		return result.UserNotExist, u
	}
	return result.UserExist, u
}


func (ur *UserRepo) Insert(user User) (code int, message *string) {
	user.Password = utils.Encrypt(user.Password)
	err := DB.Create(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *UserRepo) DeleteVia(userID uint) (code int, message *string) {
	user := User{}
	user.ID = userID
	err := DB.Debug().Select(clause.Associations).Delete(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *UserRepo) Edit(user User) (code int, message *string) {
	err := DB.Model(&user).Select("username", "avatar", "role").Updates(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *UserRepo) ChangePassword(user User) (code int, message *string) {
	user.Password = utils.Encrypt(user.Password)
	err := DB.Model(&user).Select("password").Updates(&user).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (ur *UserRepo) GetVia(userID int) (code int, message *string, user APIUser) {
	var apiUser APIUser
	apiUser.ID = uint(userID)
	err := DB.Model(&User{}).First(&apiUser).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, apiUser
	}
	return result.Success, nil, apiUser
}
func (ur *UserRepo)GetAllUsers(username string, pageSize int, pageNumber int) (int, *string, []APIUser, int64) {
	var users []APIUser
	var err error
	var total int64
	if username == "" {//select all users
		DB.Model(&[]User{}).Count(&total)
		if err = DB.Model(&[]User{}).Limit(pageSize).Offset(pageNumber).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	} else { // select all users with the same username
		DB.Model(&[]User{}).Where("username like ?", username).Count(&total)
		if err = DB.Model(&[]User{}).Where("username like ?", username).Limit(pageSize).Offset(pageNumber).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	}
	return result.Success, nil, users, total
}
func (ur *UserRepo) Login(email string, password string) (code int, message *string) {
	user := User{Email: email}

	if err := DB.Select("password").First(&user).Error; err != nil {
		return result.UserNotExist, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return result.UserPasswordNotRight, nil
	}
	return result.Success, nil
}
func (ur *UserRepo) DeleteAll() {
	DB.Exec("DELETE FROM users")
}