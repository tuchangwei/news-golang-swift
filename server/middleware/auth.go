package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/utils/result"
	"server/utils/settings"
	"strings"
	"time"
)
//Doc: https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible
var JwtKey = []byte(settings.JWTKey)
type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
//generate token
func GenerateToken(email string) (string, error)  {
	claims := MyClaims{
		Email: email ,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*15).Unix(),
			Issuer: "go-news",

		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
//verify token
func verifyToken(tokenStr string) (code int, message *string, email *string)  {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return result.TokenMalformed, nil, nil
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return result.TokenExpired, nil, nil
		} else {
			msg := ve.Error()
			return result.TokenInvalided, &msg, nil
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return result.Success, nil, &claims.Email
		}
	}
	return result.TokenInvalided, nil, nil
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
		code, msg, email := verifyToken(tokenStr)
		if code != result.Success {
			c.JSON(http.StatusOK, result.CodeMessage(code, msg))
			c.Abort()
			return
		}
		userRepo := db.UserRepo{}
		_, user := userRepo.CheckExistViaEmail(*email)
		c.Set("kCurrentUser", *user)
		c.Next()
	}
}