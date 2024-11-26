package router

import (
	"github.com/gin-gonic/gin"
	"SimpleGolang/controller"
)

func MainRouter(route *gin.Engine) {
	// Simple Application Router
	route.GET("/", controller.MainFile)
	route.GET("/GetAllProduct", controller.GetAllProduct)
	route.GET("/GetProduct/:ProductID", controller.GetProductID)
}