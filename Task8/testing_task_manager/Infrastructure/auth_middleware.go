package infrastructure

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService JWTService
}

func NewAuthMiddleware(jwtService JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (a *AuthMiddleware) AuthMiddleware(onlyAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		tokenString := authParts[1]
		claims, err := a.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		role := claims["role"].(string)
		exp := int64(claims["exp"].(float64))

		if time.Now().Unix() > exp {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		
		if role == "User" && onlyAdmin {
			c.JSON(403, gin.H{"error": "User role not allowed to access this endpoint"})
			c.Abort()
			return
		}

		c.Next()
	}
}
