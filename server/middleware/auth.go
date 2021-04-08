package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/utils/result"
	"server/utils/settings"
	"strings"
	"time"
)
//Doc: https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible
var JWT_Key = []byte(settings.JWTKey)
type MyClaims struct {
	email string `json:"email"`
	jwt.StandardClaims
}
//generate token
func GenerateToken(email string) (string, error)  {
	claims := MyClaims{
		email:          email ,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*15).Unix(),
			Issuer: "go-news",

		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_Key)
}
//verify token
func verifyToken(tokenStr string) (code int, message *string)  {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_Key, nil
	})
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return result.TokenMalformed, nil
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return result.TokenExpired, nil
		} else {
			msg := ve.Error()
			return result.TokenInvalided, &msg
		}
	}
	if token != nil {
		if _, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return result.Success, nil
		}
	}
	return result.TokenInvalided, nil
}
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, result.CodeMessage(result.NoToken, nil))
			c.Abort()
			return
		}
		strs := strings.Split(authorization, " ")
		if len(strs) != 2 || strs[0] != "Bearer" {
			c.JSON(http.StatusOK, result.CodeMessage(result.TokenFormatNotRight, nil))
			c.Abort()
			return
		}
		tokenStr := strs[1]
		if code, msg := verifyToken(tokenStr); code != result.Success {
			c.JSON(http.StatusOK, result.CodeMessage(code, msg))
			c.Abort()
			return
		}
		c.Next()
	}
}