package controllers

import (
	"net/http"

	"github.com/DevBurger/initializers"
	"github.com/DevBurger/models"
	"github.com/gin-gonic/gin"
)

// ListOrders retorna todos os pedidos
func ListOrders(c *gin.Context) {
	var orders []models.Order

	if err := initializers.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrder retorna um pedido específico pelo ID
func GetOrder(c *gin.Context) {
	id := c.Param("id")

	var order models.Order
	if err := initializers.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	c.JSON(http.StatusOK, order)
}
