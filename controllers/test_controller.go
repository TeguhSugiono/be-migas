package controllers

import (
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {

	c.JSON(200, gin.H{
		"success": true,
		"user_id": c.GetInt("user_id"),
		"role":    c.GetString("role"),
	})
}
