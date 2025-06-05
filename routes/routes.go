package routes

import (
	"DevBurger/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Rotas de pedidos
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", controllers.CreateOrder)
		orderRoutes.GET("/", controllers.GetOrders)
	}

	// Rotas de produtos
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.GET("/", controllers.GetProducts)
	}
}
