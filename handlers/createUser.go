package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/dibrinsofor/mlsa3/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lithammer/shortuuid/v4"
)

type JWTAuthDetails struct {
	UserID string
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTAuthDetails{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.FirstName,
			ExpiresAt: expiresAt,
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTOKENKEY")))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

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

	token, err := GenerateJWT(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate jwt token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully created.",
		"data":    newUser,
		"token":   token,
	})

}
