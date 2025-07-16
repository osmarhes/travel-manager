package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

// jwtKey retorna a chave secreta a partir da variável de ambiente
func jwtKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// GenerateJWT cria e assina um token JWT com duração de 24 horas
func GenerateJWT(userID uint, email string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey())
}
