package controllers

import (
	"net/http"

	"DevBurger/models"

	"github.com/gin-gonic/gin"
)

// CreateOrder cria um novo pedido
func CreateOrder(c *gin.Context) {
	var order models.Order

	// Bind JSON recebido para o struct Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Aqui normalmente salvaria no banco, mas vamos simular
	// db.Create(&order)

	c.JSON(http.StatusCreated, gin.H{"message": "Pedido criado com sucesso", "order": order})
}

// GetOrders retorna a lista de pedidos (simulação)
func GetOrders(c *gin.Context) {
	// Aqui você buscaria do banco, mas vamos simular uma lista
	orders := []models.Order{
		{ID: 1, CustomerName: "João", Total: 50.0},
		{ID: 2, CustomerName: "Maria", Total: 75.5},
	}

	c.JSON(http.StatusOK, orders)
}
