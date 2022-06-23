package handlers

import (
	"net/http"
	"time"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/dibrinsofor/mlsa3/queues"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	// azure: https://www.youtube.com/watch?v=Te9bF01iqWM

	// err := c.BindJSON(&newUser)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "failed to parse user request. check documentation.",
	// 	})
	// 	return
	// }

	c.Request.ParseMultipartForm(1000)

	// if newUser.Email == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "must provide a valid email",
	// 	})
	// 	return
	// }
	email := c.PostForm("email")
	if email != "" {
		newUser.Email = email
	}

	first := c.PostForm("first")
	if first != "" {
		newUser.FirstName = first
	}

	last := c.PostForm("last")
	if last != "" {
		newUser.LastName = last
	}

	newUser.ID = shortuuid.New()
	newUser.CreatedAt = time.Now()
	queues.SendMessage(newUser.Email)

	// _, err := redis.AddUserInstance(&newUser)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to write user data."})
	// 	return
	// }

	// token, err := middlewares.GenerateJWT(&newUser)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate jwt token"})
	// 	return
	// }

	// c.Request.Header.Set("Authorization", ("Bearer " + token))

	c.HTML(http.StatusAccepted, "success.tmpl", gin.H{
		"title":      "Message Queues Demo",
		"subheading": "Account Successfully Created",
		"first":      newUser.FirstName,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "user successfully created.",
	// 	"data":    newUser,
	// 	"token":   token,
	// })

}
