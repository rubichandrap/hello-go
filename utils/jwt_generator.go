package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken() (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_TTL"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
