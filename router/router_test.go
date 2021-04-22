package router

import (
	bytes "bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	v1 "server/api/v1"
	"server/db"
	"server/utils/result"
	"server/utils/settings"
	"testing"
)
type ResponseData struct {
	Result int `json:"result"`
	Message string `json:"message"`
}
var router *Router
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
func setup() {
	fmt.Println("_________test start__________")
	settings.InitSettings()
	db.InitDB()
	db.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&db.User{})
	router = NewRouter()
}
func shutdown() {
	fmt.Println("_________test end__________")
}
func TestRegisterUser(t *testing.T) {
	w := httptest.NewRecorder()
	user := db.User{Email: "1@1.com", Password: "123456"}
	byteArr, err := json.Marshal(user)
	if err != nil {
		t.Fatal("error:", err)
	}
	req, _ := http.NewRequest(http.MethodPost,
		router.NormalRouter.BasePath() + "/register",
		bytes.NewBuffer(byteArr))
	router.Engine.ServeHTTP(w, req)
	type ResponseData struct {
		Result int `json:"result"`
		Message string `json:"message"`
	}
	assert.Equal(t, 200, w.Code)

	var data = ResponseData{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
}
func TestLoginUser(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	gotCode, _ := user.Insert()
	assert.Equal(t, gotCode, result.Success)
	login(email, password, t)
}
func TestGetUsers(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	user.Insert()

	for i := 0; i < 20; i++ {
		user := db.User{
			Email: fmt.Sprintf("test-%d@%d.com", i,i),
			Password: "123456",
		}
		user.Insert()
	}
	token := login(email, password, t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, router.AuthorizedRouter.BasePath() + "/users",nil)
	req.Header.Add("Authorization", "Bearer " + token)
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	type ResponseDataWithUsers struct {
		ResponseData
		Data []db.User `json:"data"`
		Total int `json:"total"`
	}
	var data = ResponseDataWithUsers{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
	assert.Equal(t, 21, data.Total)
	t.Log(w.Body.String())
}
func TestDeleteUser(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	user.Insert()

	token := login(email, password, t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("%s/users/%d",router.AuthorizedRouter.BasePath() , user.ID),
		nil)
	req.Header.Add("Authorization", "Bearer " + token)
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var data = ResponseData{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
	t.Log(w.Body.String())
}

func TestEditUser(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	user.Insert()

	token := login(email, password, t)

	modifiedUser := db.User{Email: email, Password: password}
	modifiedUser.Avatar = "xxx_xxx"
	modifiedUser.Role = 2
	modifiedUser.Username = "adjkalfjdla"
	bodyBytes, _ := json.Marshal(modifiedUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/users/%d",router.AuthorizedRouter.BasePath() , user.ID),
		bytes.NewBuffer(bodyBytes))
	req.Header.Add("Authorization", "Bearer " + token)
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var data = ResponseData{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)

	_, _, apiUser := user.GetViaID()
	assert.Equal(t, modifiedUser.Role, apiUser.Role)
	assert.Equal(t, modifiedUser.Avatar, apiUser.Avatar)
	assert.Equal(t, modifiedUser.Username, apiUser.Username)
}

func TestGetUser(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	user.Insert()

	token := login(email, password, t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/users/%d",router.AuthorizedRouter.BasePath() , user.ID),
		nil)
	req.Header.Add("Authorization", "Bearer " + token)
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	type ResponseDataWithUser struct {
		ResponseData
		Data db.User `json:"data"`
	}
	var data = ResponseDataWithUser{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
	t.Log(w.Body.String())
}
func TestChangePassword(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	user.Insert()

	token := login(email, password, t)

	w := httptest.NewRecorder()
	newPwd := v1.Password{Password: "12345678"}
	pwdBytes, _ := json.Marshal(newPwd)
	req, _ := http.NewRequest(http.MethodPost, router.AuthorizedRouter.BasePath() + "/changePassword", bytes.NewBuffer(pwdBytes))
	req.Header.Add("Authorization", "Bearer " + token)
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	type ResponseDataWithUser struct {
		ResponseData
		Data db.User `json:"data"`
	}
	var data = ResponseDataWithUser{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
	token = login(email, newPwd.Password, t)
}
func TestCreatePost(t *testing.T) {
	email := "1@1.com"
	password := "123456"
	user := &db.User{Email: email, Password: password}
	user.Insert()

	token := login(email, password, t)
	post := db.Post{
		Title:    "Hello",
		PostType: 0,
		Content:  "This is a post",
	}
	postBytes, _ := json.Marshal(post)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, router.AuthorizedRouter.BasePath() + "/posts", bytes.NewBuffer(postBytes))
	req.Header.Add("Authorization", "Bearer " + token)
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var data = ResponseData{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
}
func TestDeletePost(t *testing.T) {

}

func login(email string, password string, t *testing.T) (token string) {
	t.Helper()
	w := httptest.NewRecorder()
	loginInfo := db.User{Email:email , Password: password}
	byteArr, err := json.Marshal(loginInfo)
	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost,
		router.NormalRouter.BasePath() + "/login",
		bytes.NewBuffer(byteArr))
	router.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	type ResponseDataWithToken struct {
		ResponseData
		Token string `json:"token"`
	}
	var data = ResponseDataWithToken{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, result.GetMessage(result.Success), data.Message)
	t.Log(w.Body.String())
	return data.Token
}
