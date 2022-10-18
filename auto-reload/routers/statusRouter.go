package routers

import (
	"auto-reload/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine  {
	router := gin.Default()

	router.GET("/", controllers.GetAllStatus)

	return router
}