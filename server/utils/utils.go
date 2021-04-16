package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"server/utils/result"
)

func HandleBindJSON(data interface{}, c *gin.Context) error {
	if err := c.ShouldBindJSON(&data); err != nil {
		msg := "Bind json error:" + err.Error()
		c.JSON(http.StatusOK, result.CodeMessage(result.Error, &msg))
		c.Abort()
		return err
	}
	return nil
}
func Encrypt(password string) (encryptedPassword string)  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes)
}
