package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dibrinsofor/mlsa3/middlewares"
	"github.com/dibrinsofor/mlsa3/models"
	"github.com/dibrinsofor/mlsa3/redis"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	// check if jwt exists and redirect
	// azure: https://www.youtube.com/watch?v=Te9bF01iqWM

	err := c.BindJSON(&newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation.",
		})
		return
	}

	if newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide a valid email. field cannot be left empty.",
		})
		return
	}

	newUser.ID = shortuuid.New()
	newUser.CreatedAt = time.Now()

	_, err = redis.AddUserInstance(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to write user data."})
		return
	}

	token, err := middlewares.GenerateJWT(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate jwt token"})
		return
	}
	c.Request.Header.Set("Authorization", ("Bearer " + token))

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully created.",
		"data":    newUser,
		"token":   token,
	})

}
