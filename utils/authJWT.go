package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var ld = godotenv.Load(".env")
var secret = os.Getenv("JWT_SECRET")
var secretKey = []byte(secret)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// JWT middleware logic
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token missing"})
			c.Abort()
			return
		}

		token, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		userID := claims["user_id"].(string)
		log.Print(userID)
		c.Set("user_id", userID)
		c.Next()

	}
}

func GenerateToken(id string, secret string) (string, error) {
	// Token generation logic
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	token, err := t.SignedString([]byte(secret))
	return token, err
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	return token, err
}
