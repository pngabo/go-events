package main

import (
	"events-mgmt-portal/db"
	"events-mgmt-portal/middlewares"
	"events-mgmt-portal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	// Apply middleware
	server.Use(middlewares.LoggingMiddleware())
	server.Use(middlewares.CORSMiddleware())
	server.Use(middlewares.XSSProtectionMiddleware())

	routes.RegisterRoutes(server)

	// ping server
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.Run(":3000")
}
