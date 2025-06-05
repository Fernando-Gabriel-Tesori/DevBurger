package routes

import (
	"github.com/DevBurger/controllers"
	"github.com/DevBurger/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra todas as rotas da aplicação.
func RegisterRoutes(router *gin.Engine) {
	// Rotas públicas
	router.POST("/signup", controllers.Register)
	router.POST("/login", controllers.Login)

	// Grupo de rotas protegidas por autenticação
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// Rotas de usuários
	protected.GET("/users", controllers.GetUsers)
	protected.GET("/users/:id", controllers.GetUserByID)
	protected.PUT("/users/:id", controllers.UpdateUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)

	// Rotas de produtos
	protected.POST("/products", controllers.CreateProduct)
	protected.GET("/products", controllers.GetProducts)
	protected.GET("/products/:id", controllers.GetProductByID)
	protected.PUT("/products/:id", controllers.UpdateProduct)
	protected.DELETE("/products/:id", controllers.DeleteProduct)

	// Rotas de pedidos
	protected.POST("/orders", controllers.CreateOrder)
	protected.GET("/orders", controllers.GetOrders)
	protected.GET("/orders/:id", controllers.GetOrderByID)
}
