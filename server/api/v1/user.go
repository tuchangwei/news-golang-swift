package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/utils"
	"server/utils/result"
	"server/utils/validator"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user db.User
	if utils.HandleBindJSON(user, c) != nil {
		return
	}
	var code int
	var msg *string
	code, msg = validator.Validate(user)
	if code != result.Success {
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
	var user db.User
	if utils.HandleBindJSON(&user, c) != nil {
		return
	}

	user.ID = uint(id)
	var code int
	var msg *string
	code = user.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	code, msg = user.Edit()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func ChangeUserPassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user db.User
	if utils.HandleBindJSON(&user, c) != nil {
		return
	}
	user.ID = uint(id)
	var code int
	var msg *string
	code = user.CheckExistViaID()
	if code == result.UserNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
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
	var code int
	var msg *string
	code, msg = user.Get()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, msg)
	codeMsg["data"] = user
	c.JSON(http.StatusOK, codeMsg)
}
