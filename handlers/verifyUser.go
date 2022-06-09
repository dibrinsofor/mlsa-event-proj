package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dibrinsofor/mlsa3/middlewares"
	"github.com/dibrinsofor/mlsa3/models"
	"github.com/gin-gonic/gin"
)

func VerifyUser(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		// redirect to create uri
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, claims, err := middlewares.VerifyJWT(tokenString)
	if err != nil {
		log.Println(err)
		// redirect to create uri
		return
	}

	if !token.Valid {
		// redirect to uri
		return
	}

	c.Set("user_id", claims.UserID)

	var user models.User
	// search with id and marshall redis response into user object

	c.JSON(http.StatusOK, gin.H{
		"firstname": user.FirstName,
	})

}
