package services

import (
	"errors"

	"github.com/DevBurger/database"
	"github.com/DevBurger/models"
	"github.com/DevBurger/utils"
	"gorm.io/gorm"
)

// CreateUser cria um novo usuário com hash de senha.
func CreateUser(user *models.User) error {
	// Verifica se e-mail já existe
	var existing models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return errors.New("e-mail já cadastrado")
	}

	// Hash da senha
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return database.DB.Create(user).Error
}

// AuthenticateUser valida credenciais e retorna o usuário autenticado.
func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("senha inválida")
	}

	return &user, nil
}

// GetAllUsers retorna todos os usuários cadastrados.
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retorna um usuário pelo ID.
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser atualiza os dados de um usuário.
func UpdateUser(id uint, updatedData models.User) error {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}

	// Atualiza os campos permitidos
	user.Name = updatedData.Name
	user.Email = updatedData.Email

	return database.DB.Save(&user).Error
}

// DeleteUser remove um usuário do banco.
func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
