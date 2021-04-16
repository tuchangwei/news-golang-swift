package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	v1 "server/api/v1"
	"server/middleware"
	"server/utils/settings"
)

const (
	Version = "v1"
)

func InitRouter() {
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

	baseURL := fmt.Sprintf("api/%s", Version)
	authorizedRouter := engine.Group(baseURL)
	authorizedRouter.Use(middleware.VerifyToken())
	userHandler := v1.NewUserHandler()
	{
		authorizedRouter.GET("users", userHandler.GetUsers)
		authorizedRouter.DELETE("users/:id", userHandler.DeleteUser)
		authorizedRouter.PUT("users/:id", userHandler.EditUser)
		authorizedRouter.GET("users/:id", userHandler.GetUser)
		authorizedRouter.POST("changPassword", userHandler.ChangeUserPassword)
	}

	normalRouter := engine.Group(baseURL)
	{
		normalRouter.POST("login", userHandler.Login)
		normalRouter.POST("register", userHandler.CreateUser)
	}

	err := engine.Run(":"+ settings.HttpPort)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}