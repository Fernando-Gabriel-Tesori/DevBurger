package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT cria um token JWT com ID do usuário e expiração.
func GenerateJWT(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET não definido nas variáveis de ambiente")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expira em 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateJWT valida e decodifica um token JWT.
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET não definido nas variáveis de ambiente")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Valida o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return token, nil
}

// ExtractUserID extrai o user_id do token JWT válido.
func ExtractUserID(token *jwt.Token) (uint, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if uid, ok := claims["user_id"].(float64); ok {
			return uint(uid), nil
		}
	}
	return 0, errors.New("não foi possível extrair o user_id do token")
}
