package controllers

import (
	"net/http"
	"strconv"

	"github.com/DevBurger/database"
	"github.com/DevBurger/initializers"
	"github.com/DevBurger/models"
	"github.com/gin-gonic/gin"
)

// GetUsers retorna todos os usuários cadastrados
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser retorna um usuário pelo ID
func GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash da senha (supondo função HashPassword no pacote utils)
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar senha"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// UpdateUser atualiza os dados de um usuário
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email" binding:"omitempty,email"`
		Password string `json:"password" binding:"omitempty,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar senha"})
			return
		}
		user.Password = hashedPassword
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// DeleteUser remove um usuário pelo ID
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result := initializers.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Função para hash de senha (exemplo simples, deve estar no utils)
func HashPassword(password string) (string, error) {
	// Implemente aqui seu hash de senha, exemplo com bcrypt
	// return utils.HashPassword(password)
	return password, nil // placeholder, substitua pela implementação real
}
