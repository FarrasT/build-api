package routers

import (
	"build-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/api/order", controllers.AddOrder)
	router.GET("/api/order", controllers.GetAllData)
	router.PUT("/api/order/:orderId", controllers.UpdateOrder)
	router.DELETE("/api/order/:orderId", controllers.DeleteOrder)
	router.DELETE("/api/order/:orderId/item/:itemId", controllers.DeleteItem)

	return router
}
