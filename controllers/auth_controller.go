package controllers

import (
	"BackendEsp32/models"
	"BackendEsp32/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	query := `
	INSERT INTO users (
		full_name,
		email,
		password
	)
	VALUES (?, ?, ?)
	`

	_, err := db.Exec(
		query,
		input.FullName,
		input.Email,
		string(hashPassword),
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Register berhasil",
	})
}

func Login(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	var user models.User

	query := `
	SELECT
		id,
		full_name,
		email,
		password,
		role
	FROM users
	WHERE email = ?
	LIMIT 1
	`

	err := db.QueryRow(
		query,
		input.Email,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Email tidak ditemukan",
		})

		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Password salah",
		})

		return
	}

	token, _ := utils.GenerateJWT(
		user.ID,
		user.Role,
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login berhasil",
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}
