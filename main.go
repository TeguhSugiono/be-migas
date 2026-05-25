package main

import (
	"BackendEsp32/config"
	"BackendEsp32/connection"
	"BackendEsp32/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.InitEnv()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	db := connection.SetupConnection()

	api := r.Group("/migas/api/v1")

	api.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.ApiRoutes(api)

	port := config.GetEnv("APP_PORT")

	r.Run(":" + port)
}
