package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(onlyAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT-SECRET")),nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.IndentedJSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		} 

		role := claims["role"].(string)
		exp := int64(claims["exp"].(float64))

		if time.Now().Unix() > exp {
			c.IndentedJSON(401, gin.H{"error": "token expired"})
			c.Abort()
			return
		}
		
		if role == "User" && onlyAdmin {
			c.IndentedJSON(403,gin.H{"error": "user role not allowed to access this end point"})
			c.Abort()
			return
		}

		c.Next()
	}
}
