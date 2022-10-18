package routers

import (
	"assignment-two/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrdersWithItems)
	router.DELETE("/orders/:orderId", controllers.DeleteOrderById)
	router.PUT("/orders/:orderId", controllers.UpdateOrderById)

	return router
}