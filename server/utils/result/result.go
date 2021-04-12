package result

import "github.com/gin-gonic/gin"

const (
	CodeKey    = "result"
	MessageKey = "message"

	Success = 1
	Error   = 0

	//code = 1000... user error
	UserExist = 1001
	UserNotExist = 1002
	UserRoleValueNotRight = 1003
	UserPasswordNotRight = 1004
	CantGenerateToken = 1005
	UserHasNoPermission = 1007

	//code = 2000...token error
	TokenMalformed = 2000 // Token is malformed
	TokenInvalided = 2001             // Token could not be verified because of signing problems
	TokenExpired = 2002                  // Signature validation failed
	NoToken = 2003
	TokenFormatNotRight = 2004
)
var resultCodeMsg = map[int]string{
	Success: "OK",
	Error:   "Failed",
	UserExist: "User has existed",
	UserNotExist: "User doesn't exist",
	UserRoleValueNotRight: "User's role only can be assigned to 1 or 2",
	UserPasswordNotRight: "Password is incorrect",
	CantGenerateToken: "Can't generate token",
	TokenMalformed: "Token is malformed",
	TokenInvalided: "Token is invalided",
	TokenExpired: "Token is expired",
	NoToken: "No token",
	TokenFormatNotRight: "Token format is incorrect",
	UserHasNoPermission: "User has no permission",
}

func CodeMessage(resultCode int, message *string) gin.H {
	var msg string
	if message == nil {
		msg = resultCodeMsg[resultCode]
	} else {
		msg = *message
	}
	return gin.H {
		CodeKey:    resultCode,
		MessageKey: msg,
	}
}