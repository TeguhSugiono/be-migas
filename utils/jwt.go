package utils

import (
	"BackendEsp32/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(
	userID int,
	role string,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp": time.Now().Add(
			time.Hour * 24 * 30,
		).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(config.GetEnv("JWT_SECRET")),
	)
}
