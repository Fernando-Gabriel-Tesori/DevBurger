package controllers

import (
	"net/http"
	"strconv"

	"github.com/DevBurger/database"
	"github.com/DevBurger/models"
	"github.com/gin-gonic/gin"
)

// GetProductByID busca um produto específico pelo ID
func GetProductByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, product)
}
