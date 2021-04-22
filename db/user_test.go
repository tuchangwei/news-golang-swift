package db

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/utils/result"
	"testing"
)


func InsertUser(t *testing.T) User {
	user := &User{Email: "changweitu@gmail.com", Password: "123456"}
	gotCode, gotMsg := user.Insert()
	if gotCode != result.Success {
		t.Fatal("insert failed, msg: ", *gotMsg)
	}
	return *user
}
func TestUser_CheckExistViaEmail(t *testing.T) {
	user := InsertUser(t)
	user.ID = 2
	gotCode := user.CheckExistViaEmail()
	if gotCode != result.UserExist {
		t.Fatal("got code:", gotCode)
	}
}
func TestUser_CheckExistViaID(t *testing.T) {
	user := InsertUser(t)
	gotCode := user.CheckExistViaID()
	if gotCode != result.UserExist {
		t.Fatal("got code:", gotCode)
	}
}
func TestUser_Insert(t *testing.T) {
	InsertUser(t)
}
func TestUser_DeleteViaID(t *testing.T) {
	user := InsertUser(t)
	user.DeleteViaID()
}
func TestUser_EditAndGet(t *testing.T) {
	user := InsertUser(t)
	user.Username = "tu"
	user.Avatar = "xxx___xxx"
	user.Role = 2
	user.Edit()
	user1 := User{}
	user1.ID = user.ID
	_, _, apiUser := user1.GetViaID()
	if apiUser.Username != user.Username || apiUser.Avatar != user.Avatar || apiUser.Role != user.Role {
		t.Fatalf("%+v \n %+v", user, user1)
	}
}
func TestUser_ChangePassword(t *testing.T) {
	pwd := "666666"
	user := InsertUser(t)
	t.Log(user.Password)
	user.Password = pwd
	code, msg := user.ChangePassword()
	if code != result.Success {
		t.Fatal(*msg)
	}
	gotUser := User{
		Email: user.Email,
	}
	gotUser.CheckExistViaEmail()
	err := bcrypt.CompareHashAndPassword([]byte(gotUser.Password), []byte(pwd))
	if err != nil {
		t.Fatal("Change Password Error:", err)
	}
}
func TestUser_GetAllUsers(t *testing.T) {
	user := User{Email: "1@1.com", Password: "123456"}
	user.Insert()
	for i := 0; i <20; i++ {
		user.Email = fmt.Sprintf("test%d@%d.com", i, i)
		user.Insert()
	}
	code, msg, users, total := user.GetAllUsers("", 10, 3)
	if code != result.Success {
		t.Fatal(*msg)
	}
	t.Logf("users:%+v, users' count:%d, total count: %d", users, len(users), total)
}
func TestUser_LoginWithEmailAndPassword(t *testing.T) {
	email := "1@1.com"
	pwd := "123456"
	user := User{Email: email, Password: pwd}
	user.Insert()
	user1 := User{Email: email, Password: pwd}
	code, _ := user1.LoginWithEmailAndPassword()
	if code != result.Success {
		t.Fatalf("%+v", result.CodeMessage(code, nil))
	}
}