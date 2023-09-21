package services

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key")

func GenerateToken(userID uint) (string, error) {
	// Generate initial token
	token := jwt.New(jwt.SigningMethodHS256)

	// Define claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["expire"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign token
	t, err := token.SignedString(jwtSecret)
	return t, err
}
