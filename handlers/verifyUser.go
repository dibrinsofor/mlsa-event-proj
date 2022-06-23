package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dibrinsofor/mlsa3/middlewares"
	"github.com/gin-gonic/gin"
)

func VerifyUser(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.Redirect(http.StatusUnauthorized, "/createUser")
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, claims, err := middlewares.VerifyJWT(tokenString)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusUnauthorized, "/createUser")
		c.Abort()
		return
	}

	if !token.Valid {
		c.Redirect(http.StatusUnauthorized, "/createUser")
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserID)

	// user := redis.FindUserByID(claims.UserID)

	// c.JSON(http.StatusOK, gin.H{
	// 	"firstname": user.FirstName,
	// })

}
