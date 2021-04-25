package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/middleware"
	"server/utils"
	"server/utils/result"
	"server/utils/validator"
	"strconv"
)
type Password struct {
	Password string `json:"password"`
}
type UserHandler struct {
}
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (uh *UserHandler)CreateUser(c *gin.Context) {
	var user db.User
	if utils.HandleBindJSON(&user, c) != nil {
		return
	}
	var code int
	var msg *string
	code, msg = validator.Validate(user)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	user.CheckExistViaEmail()
	if code == result.UserExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	code, msg = user.Insert()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	token, err := middleware.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, result.CodeMessage(result.CantGenerateToken, nil))
		c.Abort()
		return
	}
	responseData := result.CodeMessage(result.Success, nil)
	responseData["token"] = token
	c.JSON(http.StatusOK, responseData)
}
func checkUserPermission(c *gin.Context) bool {
	id, _ := strconv.Atoi(c.Param("id"))
	currentUser := middleware.GetCurrentUserInContext(c)
	if currentUser.ID != uint(id) && currentUser.Role != 2 {
		c.JSON(http.StatusOK, result.CodeMessage(result.UserHasNoPermission, nil))
		c.Abort()
		return false
	}
	return true
}
func (uh *UserHandler)DeleteUser(c *gin.Context) {
	if !checkUserPermission(c) {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var code int
	var msg *string
	user := db.User{}
	user.ID = uint(id)
	code = user.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	code, msg = user.DeleteViaID()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

//EditUser edit user's role, avatar, username
func (uh *UserHandler)EditUser(c *gin.Context) {
	if !checkUserPermission(c) {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	type UserParam struct {
		Username *string `json:"username"`
		Avatar *string `json:"avatar"`
		Role *int `json:"role"`//1 normal, 2 admin
		ID   uint `json:"id"`
	}
	var param UserParam
	if utils.HandleBindJSON(&param, c) != nil {
		return
	}
	user := db.User{}
	user.ID = uint(id)
	var code int
	var msg *string
	code = user.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}

	//if param.Role is nil, that means the client didn't send us Role parameter, so we can assign it with the database value
	//if param.Role is not nil, we know the client send us the parameter, we will use the value to update database.
	if param.Role != nil {
		if *(param.Role) != 1 && *(param.Role) != 2 {
			c.JSON(http.StatusOK, result.CodeMessage(result.UserRoleValueNotRight, nil))
			c.Abort()
			return
		}
		user.Role = *param.Role
	}

	//if param.Avatar is nil, that means the client didn't send us avatar parameter, so we can assign it with the database value
	//if param.Avatar is not nil, we know the client send us the parameter, we will use the value to update database.
	if param.Avatar != nil {
		user.Avatar = *param.Avatar
	}
	//if param.Username is nil, that means the client didn't send us username parameter, so we can assign it with the database value
	//if param.Username is not nil, we know the client send us the parameter, we will use the value to update database.
	if param.Username != nil {
		user.Username = *param.Username
	}

	code, msg = user.Edit()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func (uh *UserHandler)ChangeUserPassword(c *gin.Context) {

	var password Password
	if utils.HandleBindJSON(&password, c) != nil {
		return
	}
	value, _ := c.Get("kCurrentUser")
	currentUser := value.(db.User)
	var code int
	var msg *string
	currentUser.Password = password.Password
	if code, msg := validator.Validate(currentUser); code == result.Error {//validate password length
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	code, msg = currentUser.ChangePassword()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func (uh *UserHandler)GetUser(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	var user db.User
	user.ID = uint(id)
	code, msg, apiUser := user.GetViaID()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, msg)
	codeMsg["data"] = apiUser
	c.JSON(http.StatusOK, codeMsg)
}
func (uh *UserHandler)GetUsers(c *gin.Context)  {
	username := c.Query("username")
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}
	user := db.User{}
	code, msg, users, total := user.GetAllUsers(username, pageSize, pageNum)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, nil)
	codeMsg["data"] = users
	codeMsg["total"] = total
	c.JSON(http.StatusOK, codeMsg)
}
func (uh *UserHandler)Login(c *gin.Context) {
	var user db.User
	if err := utils.HandleBindJSON(&user, c); err != nil {
		return
	}
	if code, msg := validator.Validate(user); code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	if code, msg := user.LoginWithEmailAndPassword(); code != result.Success {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	token, err := middleware.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, result.CodeMessage(result.CantGenerateToken, nil))
		c.Abort()
		return
	}
	responseData := result.CodeMessage(result.Success, nil)
	responseData["token"] = token
	c.JSON(http.StatusOK, responseData)
}
func (uh *UserHandler) Follow(c *gin.Context)  {
	user := db.User{}
	if err := utils.HandleBindJSON(&user, c); err != nil {
		msg := err.Error()
		c.JSON(http.StatusOK, result.CodeMessage(result.Error, &msg))
		c.Abort()
		return
	}
	code := user.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	currentUser := middleware.GetCurrentUserInContext(c)
	if currentUser.ID == user.ID {
		c.JSON(http.StatusOK, result.CodeMessage(result.UserCantFollowHimself, nil))
		c.Abort()
		return
	}
	if !currentUser.CheckRelationshipExist(user) {
		code, msg := currentUser.Follow(user)
		if code != result.Success {
			c.JSON(http.StatusOK, result.CodeMessage(result.Error, msg))
			c.Abort()
			return
		}
	}
	c.JSON(http.StatusOK, result.CodeMessage(result.Success, nil))
}
func (uh *UserHandler) Unfollow(c *gin.Context)  {
	user := db.User{}
	if err := utils.HandleBindJSON(&user, c); err != nil {
		msg := err.Error()
		c.JSON(http.StatusOK, result.CodeMessage(result.Error, &msg))
		c.Abort()
		return
	}
	code := user.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	currentUser := middleware.GetCurrentUserInContext(c)
	if currentUser.ID == user.ID {
		c.JSON(http.StatusOK, result.CodeMessage(result.UserCantUnfollowHimself, nil))
		c.Abort()
		return
	}
	code, msg := currentUser.Unfollow(user)
	if code != result.Success {
		c.JSON(http.StatusOK, result.CodeMessage(result.Error, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(result.Success, nil))
}
func (uh *UserHandler) GetFollowers(c *gin.Context)  {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}
	user := middleware.GetCurrentUserInContext(c)
	code, msg, users, total := user.GetFollowers(pageSize, pageNum)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, nil)
	codeMsg["data"] = users
	codeMsg["total"] = total
	c.JSON(http.StatusOK, codeMsg)
}
func (uh *UserHandler) GetFollowings(c *gin.Context)  {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}
	user := middleware.GetCurrentUserInContext(c)
	code, msg, users, total := user.GetFollowings(pageSize, pageNum)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, nil)
	codeMsg["data"] = users
	codeMsg["total"] = total
	c.JSON(http.StatusOK, codeMsg)
}
