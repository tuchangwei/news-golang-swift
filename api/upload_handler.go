package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/upload"
	"server/utils/result"
)

type UploadHandler struct {

}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}
func (u *UploadHandler) UploadPhoto(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		msg := err.Error()
		c.JSON(http.StatusOK, result.CodeMessage(result.UploadError, &msg))
		c.Abort()
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		msg := err.Error()
		c.JSON(http.StatusOK, result.CodeMessage(result.UploadError, &msg))
		c.Abort()
		return
	}
	location, err := upload.Upload(file)
	if err != nil {
		msg := err.Error()
		c.JSON(http.StatusOK, result.CodeMessage(result.UploadError, &msg))
		c.Abort()
		return
	}
	data := result.CodeMessage(result.Success, nil)
	data["url"] = location
	c.JSON(http.StatusOK, data)

}
