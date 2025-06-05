package main

import (
	"DevBurger/database"
	"DevBurger/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.Connect()            // conecta ao banco
	routes.RegisterRoutes(router) // registra as rotas

	router.Run(":3000") // roda servidor na porta 3000
}
