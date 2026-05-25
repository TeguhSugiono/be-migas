package middleware

import (
	"BackendEsp32/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Token tidak ditemukan",
			})

			c.Abort()

			return
		}

		tokenString := strings.Replace(
			authHeader,
			"Bearer ",
			"",
			1,
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				return []byte(
					config.GetEnv("JWT_SECRET"),
				), nil
			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Token tidak valid",
			})

			c.Abort()

			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set(
			"user_id",
			int(claims["user_id"].(float64)),
		)

		c.Set(
			"role",
			claims["role"].(string),
		)

		c.Next()
	}
}
