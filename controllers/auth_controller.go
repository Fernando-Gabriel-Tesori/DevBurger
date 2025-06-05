package controllers

import (
	"net/http"
	"time"

	"github.com/DevBurger/initializers"
	"github.com/DevBurger/models"
	"github.com/DevBurger/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register cria um novo usuário
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := initializers.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Usuário já cadastrado com esse email"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar senha"})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Name:     input.Name,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

// Login autentica o usuário e gera o token JWT
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Corrigido: apenas user.ID como argumento
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar token"})
		return
	}

	c.SetCookie("Authorization", token, int(time.Hour.Seconds()*24), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
	})
}

// Logout limpa o cookie do token JWT
func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}
