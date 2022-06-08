package handlers

import (
	"log"
	"net/http"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	// check if jwt exists and redirect

	err := c.BindJSON(&newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation.",
		})
		return
	}

	// generate user id uuid
	// newUser.ID =

	// set created at time

	// persist in memory

}
