package main

import (
	"api/configs"
	"api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDB()
	configs.ConnectRedis()
}

func main() {
	router := gin.Default()
	router.Static("/public", "./public") // static file

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"success": true,
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	routes.UserRoute(router) // user router

	router.Run(os.Getenv("PORT"))
}