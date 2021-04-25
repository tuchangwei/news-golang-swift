package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/utils"
	"server/utils/result"
	"time"
)

type User struct {
	gorm.Model
	Email string    `gorm:"not null" validate:"required,email" json:"email"`
	Password string `gorm:"not null" validate:"required,min=6,max=120" json:"password"`
	Posts []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserInfo
}
type Friend struct {
	FollowerID uint `gorm:"primary_key"`
	FollowingID uint `gorm:"primary_key"`
	CreatedAt time.Time
}
type UserInfo struct {
	Username string    `json:"username"`
	Avatar string      `json:"avatar"`
	Role int           `gorm:"not null;default:1" json:"role"`//1 normal, 2 admin
}
type APIUser struct {
	UserInfo
	ID   uint `json:"id"`
}

// BeforeDelete Hook, implement BeforeDeleteInterface interface
func (u *User) BeforeDelete(db *gorm.DB) error {
	db.Where("follower_id=? OR following_id=?", u.ID, u.ID).Delete(Friend{})
	return nil
}

func (u *User) CheckExistViaEmail() (code int) {
	//we set the primary key id to 0, so gorm ignore the id to build query conditions
	//If not,the query conditions will to build including email and id
	//That is not we want.
	u.ID = 0
	err := DB.Debug().Where("email=?", u.Email).First(u).Error
	if err != nil {
		return result.UserNotExist
	}
	return result.UserExist
}
func (u *User) CheckExistViaID() (code int) {
	err := DB.Where("id=?", u.ID).First(u).Error
	if err != nil {
		return result.UserNotExist
	}
	return result.UserExist
}

func (u *User) Insert() (code int, message *string) {
	u.ID = 0
	u.Password = utils.Encrypt(u.Password)
	err := DB.Create(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) DeleteViaID() (code int, message *string) {
	err := DB.Debug().Select(clause.Associations).Delete(u, u.ID).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) Edit() (code int, message *string) {
	err := DB.Select("username", "avatar", "role").Updates(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) ChangePassword() (code int, message *string) {
	u.Password = utils.Encrypt(u.Password)
	err := DB.Select("password").Updates(u).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (u *User) GetViaID() (code int, message *string, user APIUser) {
	var apiUser APIUser
	apiUser.ID = u.ID
	err := DB.Model(u).Where("id=?", u.ID).First(&apiUser).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, apiUser
	}
	return result.Success, nil, apiUser
}
func (u *User)GetAllUsers(username string, pageSize int, pageNumber int) (int, *string, []APIUser, int64) {
	var users []APIUser
	var err error
	var total int64
	if username == "" {//select all users
		DB.Model(&User{}).Count(&total)
		if err = DB.Model(&User{}).Limit(pageSize).Offset(pageNumber*pageSize).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	} else { // select all users with the same username
		DB.Model(&User{}).Where("username like ?", username).Count(&total)
		if err = DB.Model(&User{}).Where("username like ?", username).Limit(pageSize).Offset(pageNumber*pageSize).Find(&users).Error;
			err != nil {
			msg := err.Error()
			return result.Error, &msg, users, 0
		}
	}
	return result.Success, nil, users, total
}
func (u *User) LoginWithEmailAndPassword() (code int, message *string) {
	var dbUser User
	if err := DB.Select("password").Where("email=?", u.Email).First(&dbUser).Error; err != nil {
		return result.UserNotExist, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)); err != nil {
		return result.UserPasswordNotRight, nil
	}
	return result.Success, nil
}
func (u *User) DeleteAll() {
	DB.Exec("DELETE FROM users")
}
func (u *User) Follow(anotherUser User) (code int, message *string)  {
	friend := Friend{
		FollowerID:  u.ID,
		FollowingID: anotherUser.ID,
		CreatedAt:   time.Now(),
	}
	if err := DB.Create(friend).Error; err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}
func (u *User) CheckRelationshipExist(anotherUser User) bool {
	err := DB.Model(&Friend{}).Where("follower_id=? AND following_id=?", u.ID, anotherUser.ID).First(&Friend{}).Error
	if err != nil {
		return false
	}
	return true
}

func (u *User) Unfollow(anotherUser User) (code int, message *string)  {
	friend := Friend{
		FollowerID:  u.ID,
		FollowingID: anotherUser.ID,
	}
	if err := DB.Delete(friend).Error; err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}




