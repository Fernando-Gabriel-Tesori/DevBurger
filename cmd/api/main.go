package main

import (
	"log"
	"os"

	"github.com/DevBurger/config"
	"github.com/DevBurger/initializers"
	"github.com/DevBurger/middlewares"
	"github.com/DevBurger/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	err = initializers.InitDB() // Chamada para função correta
	if err != nil {
		log.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}

	router := gin.Default()

	router.Use(middlewares.LoggerMiddleware()) // Middleware de log
	router.Use(middlewares.AuthMiddleware())   // Middleware de autenticação

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}
