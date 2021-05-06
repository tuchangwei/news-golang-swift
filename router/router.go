package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/api"
	"server/middleware"
	"server/utils/settings"
)


//var NewRouter *gin.Engine
//var BaseURL string



func NewRouter() *gin.Engine {
	gin.SetMode(settings.AppMode)
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/", func(c *gin.Context) {
		host := c.Request.Host
		html := fmt.Sprintf(`
		<html>
		<body>
			<h1>Welcome, please visit <a href='http://%s/api'>http://%s/api</a> to get data</h1>
		</body>
		</html>
		`, host, host)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	engine.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, "Invalid api request") })
	userHandler := api.NewUserHandler()
	postHandler := api.NewPostHandler()
	v1 := engine.Group("api/v1")
	{
		v1.POST("login", userHandler.Login)
		v1.POST("register", userHandler.CreateUser)
		v1.GET("posts", postHandler.GetAllPosts)
		auth := v1.Use(middleware.VerifyToken())
		{
			auth.GET("users", userHandler.GetUsers)
			auth.DELETE("users/:id", userHandler.DeleteUser)
			auth.PUT("users/:id", userHandler.EditUser)
			auth.GET("users/:id", userHandler.GetUser)
			auth.POST("changePassword", userHandler.ChangeUserPassword)
			auth.POST("follow", userHandler.Follow)
			auth.POST("unfollow", userHandler.Unfollow)

			auth.GET("followers", userHandler.GetFollowers)
			auth.GET("followings", userHandler.GetFollowings)

			auth.POST("posts", postHandler.CreatePost)
			auth.DELETE("posts/:id", postHandler.DeletePost)
			//Get some user's all posts
			auth.GET("users/:id/posts", postHandler.GetAllPosts)

			uploadHandler := api.NewUploadHandler()
			auth.POST("upload", uploadHandler.UploadPhoto)
		}

	}
	return engine
}