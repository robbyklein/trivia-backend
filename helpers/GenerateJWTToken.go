package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId int) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 6).Unix(),
		"iss": "Trivia Game",
		"nbf": time.Now().Unix(),
	})

	// Convert the secret key to []byte
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Sign it
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	// Return it
	return tokenString, nil
}
