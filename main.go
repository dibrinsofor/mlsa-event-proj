package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dibrinsofor/mlsa3/handlers"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck())
	r.POST("/createUser", handlers.CreateUser)
	// r.GET("/", handlers.VerifyUser)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "Message Queues Demo",
			"subheading": "Create Account",
		})
	})

	return r
}

func main() {
	r := SetupServer()

	r.LoadHTMLGlob("templates/*")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	// watch for App_offline.htm and exit the program if present
	// This allows continuous deployment on App Service as the .exe will not be
	// terminated otherwise
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.HasSuffix(event.Name, "app_offline.htm") {
					fmt.Println("Exiting due to app_offline.htm being present")
					os.Exit(0)
				}
			}
		}
	}()

	// get the current working directory and watch it
	currentDir, err := os.Getwd()
	if err := watcher.Add(currentDir); err != nil {
		fmt.Println("ERROR", err)
	}

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "6969"
	}

	r.Run("127.0.0.1:" + port)
}
