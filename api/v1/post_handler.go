package v1

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
	postRepo *db.PostRepo
}
func NewPostHandler() *PostHandler {
	return &PostHandler{postRepo: db.NewPostRepo()}
}
func (ph *PostHandler)CreatePost(c *gin.Context) {
	var post db.Post
	if utils.HandleBindJSON(&post, c) != nil {
		return
	}
	post.UserID = middleware.GetCurrentUserInContext(c).ID
	var code int
	var msg *string
	code, msg = ph.postRepo.Insert(post)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.CodeMessage(result.Success, nil))
}

func (ph *PostHandler)DeletePost(c *gin.Context) {
	if !checkUserPermission(c) {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var code int
	var msg *string
	code, _ = ph.postRepo.CheckExistViaID(id)
	if code == result.PostNotExist {
		c.JSON(http.StatusOK, result.CodeMessage(code, nil))
		c.Abort()
		return
	}
	code, msg = ph.postRepo.DeleteVia(id)
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
	code, msg, apiUser := ph.postRepo.GetVia(id)
	if code == result.Error {
		c.JSON(http.StatusOK, result.CodeMessage(code, msg))
		c.Abort()
		return
	}
	var codeMsg = result.CodeMessage(code, msg)
	codeMsg["data"] = apiUser
	c.JSON(http.StatusOK, codeMsg)
}
func (ph *PostHandler)GetAllPosts(c *gin.Context)  {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	userID, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = 20
	}

	code, msg, posts, total := ph.postRepo.GetAllPosts(userID, pageSize, pageNum)
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
