package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/middleware"
	"server/utils"
	"server/utils/result"
	"strconv"
)

type PostHandler struct {
}
func NewPostHandler() *PostHandler {
	return &PostHandler{}
}
func (ph *PostHandler)CreatePost(c *gin.Context) {
	var post db.Post
	if utils.HandleBindJSON(&post, c) != nil {
		return
	}
	post.UserID = middleware.GetCurrentUserInContext(c).ID
	var code int
	var msg *string
	code, msg = post.Insert()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(result.Success, nil))
}

func (ph *PostHandler)DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var code int
	var msg *string
	post := db.Post{}
	post.ID = uint(id)
	code = post.CheckExistViaID()
	if code == result.PostNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	currentUser := middleware.GetCurrentUserInContext(c)
	if currentUser.ID != post.UserID && currentUser.Role != 2 {
		c.JSON(http.StatusOK, result.CodeMessage(result.UserHasNoPermission, nil))
		c.Abort()
	}
	code, msg = post.DeleteViaID()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(code, msg))
}

func (ph *PostHandler)GetPost(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	var post db.Post
	post.ID = uint(id)
	code, msg := post.GetViaID()
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, msg)
	codeMsg["data"] = post
	c.JSON(http.StatusOK, codeMsg)
}
func (ph *PostHandler)GetAllPosts(c *gin.Context)  {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	userID, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = 20
	}
	post := db.Post{}
	code, msg, posts, total := post.GetAllPosts(userID, pageSize, pageNum)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, nil)
	codeMsg["data"] = posts
	codeMsg["total"] = total
	c.JSON(http.StatusOK, codeMsg)
}
