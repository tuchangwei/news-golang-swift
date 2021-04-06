package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	v1 "server/api/v1"
	"server/utils"
)

const (
	Version = "v1"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
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
	r := engine.Group(baseURL)
	{
		r.GET("users", v1.GetUsers)
		r.POST("users", v1.CreateUser)
		r.DELETE("users/:id", v1.DeleteUser)
		r.PUT("users/:id", v1.EditUser)
		r.GET("users/:id", v1.GetUser)
		r.PUT("changPassword/:id", v1.ChangeUserPassword)
	}


	err := engine.Run(":"+utils.HttpPort)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}