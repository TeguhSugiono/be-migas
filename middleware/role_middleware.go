package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(
	roles ...string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		role := c.GetString("role")

		for _, r := range roles {

			if role == r {

				c.Next()

				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Access denied",
		})

		c.Abort()
	}
}
