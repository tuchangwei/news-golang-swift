package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/middleware"
	"server/model"
	"server/utils"
	"server/utils/result"
	"server/utils/validator"
	"strconv"
)

type UserHandler struct { }
var userRepo = db.NewUserRepo()
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (uh *UserHandler)CreateUser(c *gin.Context) {
	var user model.User
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
	code, _ = userRepo.CheckExistViaEmail(*user.Email)
	if code == result.UserExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	code, msg = userRepo.Insert(user)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	token, err := middleware.GenerateToken(*(user.Email))
	if err != nil {
		c.JSON(http.StatusOK, result.CodeMessage(result.CantGenerateToken, nil))
		c.Abort()
		return
	}
	responseData := result.CodeMessage(result.Success, nil)
	responseData["token"] = token
	c.JSON(http.StatusOK, responseData)
}
func (uh *UserHandler)DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var code int
	var msg *string
	code, _ = userRepo.CheckExistViaID(id)
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	code, msg = userRepo.DeleteVia(id)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

//allow user to edit all fields except email and password.
//if user pass into email and password, the two fields will be omitted.
func (uh *UserHandler)EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	if utils.HandleBindJSON(&user, c) != nil {
		return
	}
	user.ID = uint(id)
	var code int
	var msg *string
	code, dbUser := userRepo.CheckExistViaID(id)
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}

	//if user.Role is nil, that means the client didn't send us Role parameter, so we can assign it with the database value
	//if user.Role is not nil, we know the client send us the parameter, we will use the value to update database.
	if user.Role == nil {
		user.Role = dbUser.Role
	} else if *(user.Role) != 1 && *(user.Role) != 2 {
		c.JSON(http.StatusOK, result.CodeMessage(result.UserRoleValueNotRight, nil))
		c.Abort()
		return
	}

	//if user.Avatar is nil, that means the client didn't send us avatar parameter, so we can assign it with the database value
	//if user.Avatar is not nil, we know the client send us the parameter, we will use the value to update database.
	if user.Avatar == nil {
		user.Avatar = dbUser.Avatar
	}
	//if user.Username is nil, that means the client didn't send us username parameter, so we can assign it with the database value
	//if user.Username is not nil, we know the client send us the parameter, we will use the value to update database.
	if user.Username == nil {
		user.Username = dbUser.Username
	}

	code, msg = userRepo.Edit(user)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func (uh *UserHandler)ChangeUserPassword(c *gin.Context) {
	type Password struct {
		Password string `json:"password"`
	}
	var password Password
	if utils.HandleBindJSON(&password, c) != nil {
		return
	}
	email, exist := c.Get("email")
	if !exist {
		c.JSON(http.StatusOK, result.CodeMessage(result.NoEmailInContext, nil))
		c.Abort()
		return
	}
	var code int
	var msg *string
	code, dbUser := userRepo.CheckExistViaEmail(*email.(*string))
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	dbUser.Password = &password.Password
	if code, msg := validator.Validate(dbUser); code == result.Error {//validate password length
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	code, msg = userRepo.ChangePassword(*dbUser)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func (uh *UserHandler)GetUser(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	user.ID = uint(id)
	code, msg, apiUser := userRepo.GetVia(id)
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

	code, msg, users, total := userRepo.GetAllUsers(username, pageSize, pageNum)
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
	var user model.User
	if err := utils.HandleBindJSON(&user, c); err != nil {
		return
	}
	if code, msg := validator.Validate(user); code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	if code, msg := userRepo.Login(&user); code != result.Success {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	token, err := middleware.GenerateToken(*(user.Email))
	if err != nil {
		c.JSON(http.StatusOK, result.CodeMessage(result.CantGenerateToken, nil))
		c.Abort()
		return
	}
	responseData := result.CodeMessage(result.Success, nil)
	responseData["token"] = token
	c.JSON(http.StatusOK, responseData)
}
