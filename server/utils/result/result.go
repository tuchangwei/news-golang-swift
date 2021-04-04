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
)
var resultCodeMsg = map[int]string{
	Success: "OK",
	Error:   "Failed",
	UserExist: "User has existed",
	UserNotExist: "User doesn't exist",
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