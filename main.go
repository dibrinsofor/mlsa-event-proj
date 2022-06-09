package main

import (
	"os"

	"github.com/dibrinsofor/mlsa3/handlers"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck())
	r.POST("/createUser", handlers.CreateUser)
	r.GET("/", handlers.VerifyUser)

	return r
}

func main() {
	r := SetupServer()

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "6969"
	}

	r.Run("127.0.0.1:" + port)
}
