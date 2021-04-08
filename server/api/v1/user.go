package v1

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

func CreateUser(c *gin.Context) {
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
	code = user.CheckExistViaEmail()
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
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}
func DeleteUser(c *gin.Context) {
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
	code, msg = user.Delete()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

//allow user to edit all fields except email and password.
//if user pass into email and password, the two fields will be omitted.
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userFromRequest db.User
	if utils.HandleBindJSON(&userFromRequest, c) != nil {
		return
	}
	userFromRequest.ID = uint(id)
	var userFromDB = db.User{}
	userFromDB.ID = userFromRequest.ID
	var code int
	var msg *string
	code = userFromDB.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}

	//if userFromRequest.Role is nil, that means the client didn't send us avatar parameter, so we can assign it with the database value
	//if userFromRequest.Role is not nil, we know the client send us the parameter, we will use the value to update database.
	if userFromRequest.Role == nil {
		userFromRequest.Role = userFromDB.Role
	} else if *(userFromRequest.Role) != 1 && *(userFromRequest.Role) != 2 {
		c.JSON(http.StatusOK, result.CodeMessage(result.UserRoleValueNotRight, nil))
		c.Abort()
		return
	}

	//if userFromRequest.Avatar is nil, that means the client didn't send us avatar parameter, so we can assign it with the database value
	//if userFromRequest.Avatar is not nil, we know the client send us the parameter, we will use the value to update database.
	if userFromRequest.Avatar == nil {
		userFromRequest.Avatar = userFromDB.Avatar
	}
	//if userFromRequest.Username is nil, that means the client didn't send us username parameter, so we can assign it with the database value
	//if userFromRequest.Username is not nil, we know the client send us the parameter, we will use the value to update database.
	if userFromRequest.Username == nil {
		userFromRequest.Username = userFromDB.Username
	}

	code, msg = userFromRequest.Edit()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func ChangeUserPassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tmp db.User
	if utils.HandleBindJSON(&tmp, c) != nil {
		return
	}
	password := tmp.Password
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
	user.Password = password
	if code, msg := validator.Validate(&user); code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	code, msg = user.ChangePassword()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func GetUser(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	var user db.User
	user.ID = uint(id)
	code, msg, apiUser := user.Get()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, msg)
	codeMsg["data"] = apiUser
	c.JSON(http.StatusOK, codeMsg)
}

func GetUsers(c *gin.Context)  {
	username := c.Query("username")
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}

	code, msg, users, total := db.GetAllUsers(username, pageSize, pageNum)
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
func Login(c *gin.Context) {
	var user db.User
	if err := utils.HandleBindJSON(&user, c); err != nil {
		return
	}
	if code, msg := validator.Validate(user); code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	if code, msg := user.Login(); code != result.Success {
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
