package controllers

import (
	"net/http"
	"time"

	"github.com/DevBurger/initializers"
	"github.com/DevBurger/models"
	"github.com/gin-gonic/gin"
)

// OrderItemInput define a estrutura para criação de um item de pedido
type OrderItemInput struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  uint    `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"required,min=0"`
}

// OrderInput define a estrutura para criação de um pedido
type OrderInput struct {
	UserID uint             `json:"user_id" binding:"required"`
	Items  []OrderItemInput `json:"items" binding:"required,dive,required"`
	Total  float64          `json:"total" binding:"required"`
	Status string           `json:"status" binding:"omitempty,oneof=pending completed cancelled"`
}

// CreateOrder cria um novo pedido com seus itens
func CreateOrder(c *gin.Context) {
	var input OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		UserID:    input.UserID,
		Total:     input.Total,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tx := initializers.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}

	for _, itemInput := range input.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: itemInput.ProductID,
			Quantity:  int(itemInput.Quantity), // Cast para int
			Price:     itemInput.Price,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar item do pedido"})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha na transação"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Pedido criado com sucesso",
		"order":   order,
	})
}

// GetOrders retorna todos os pedidos com seus itens
func GetOrders(c *gin.Context) {
	var orders []models.Order

	if err := initializers.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrderByID retorna um pedido específico pelo ID
func GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")

	var order models.Order
	if err := initializers.DB.Preload("Items").First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
