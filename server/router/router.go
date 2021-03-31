package router

import (
	"github.com/gin-gonic/gin"
	"server/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Run(utils.HttpPort)
}