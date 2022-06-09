package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func VerifyJWT(tokenString string) (*jwt.Token, JWTAuthDetails, error) {
	var claims JWTAuthDetails
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTOKENKEY")), nil
	})

	return token, claims, err
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no token found in authorization header"})
			return
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid format in authorization header"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, claims, err := VerifyJWT(tokenString)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "token contains an invalid number of segments"})
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", claims.UserID)

	}
}
