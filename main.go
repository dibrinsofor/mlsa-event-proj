package main

import (
	"github.com/dibrinsofor/mlsa3/handlers"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck())
	r.POST("/createUser", handlers.CreateUser)

	return r
}

func main() {
	r := SetupServer()

	r.Run(":6969")
}
