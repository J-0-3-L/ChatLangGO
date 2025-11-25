package auth

import (
	"log"
	"net/http"
	"strings"

	//"go000/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("SECRETO_PASS")

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println(">>> MIDDLEWARE EJECUTADO <<<")

		authHeader := c.GetHeader("Authorization")
		log.Println("AUTH_HEADER_RAW:", authHeader)

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println("TOKEN_STRING:", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//log.Println("ERROR_AUTH:", err)
			return Secret, nil
		})

		if err != nil || !token.Valid {
			log.Println("JWT_ERROR:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", uint(claims["user_id"].(float64)))

		c.Next()
	}
}
