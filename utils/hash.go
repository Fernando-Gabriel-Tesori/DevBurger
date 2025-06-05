package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword gera um hash seguro a partir da senha fornecida.
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("a senha não pode estar vazia")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// CheckPasswordHash compara uma senha com seu hash e retorna true se forem compatíveis.
func CheckPasswordHash(password, hash string) bool {
	if password == "" || hash == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
