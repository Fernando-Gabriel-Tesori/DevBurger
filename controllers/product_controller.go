package controllers

import (
	"net/http"
	"strconv"

	"github.com/DevBurger/initializers"
	"github.com/DevBurger/models"
	"github.com/gin-gonic/gin"
)

// GetProducts retorna todos os produtos cadastrados
func GetProducts(c *gin.Context) {
	var products []models.Product

	result := initializers.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct retorna um produto pelo ID
func GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var product models.Product
	result := initializers.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct cria um novo produto
func CreateProduct(c *gin.Context) {
	var input struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	if err := initializers.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct atualiza os dados de um produto existente
func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var product models.Product
	result := initializers.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"omitempty,gt=0"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Price > 0 {
		product.Price = input.Price
	}

	if err := initializers.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct remove um produto pelo ID
func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result := initializers.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.Status(http.StatusNoContent)
}
