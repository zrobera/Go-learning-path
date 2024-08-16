package infrastructure_test

import (
	"os"
	"testing"
	"time"

	infrastructure "test_task_manager/Infrastructure"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	jwtService := infrastructure.NewJWTService()

	username := "testuser"
	role := "User"

	token, err := jwtService.GenerateToken(username, role)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	jwtService := infrastructure.NewJWTService()

	username := "testuser"
	role := "User"

	token, err := jwtService.GenerateToken(username, role)
	assert.NoError(t, err)

	claims, err := jwtService.ValidateToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, username, claims["username"])
	assert.Equal(t, role, claims["role"])
}

func TestValidateToken_InvalidToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	jwtService := infrastructure.NewJWTService()

	invalidToken := "invalid.token.string"
	claims, err := jwtService.ValidateToken(invalidToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	jwtService := infrastructure.NewJWTService()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "testuser",
		"role":     "User",
		"exp":      time.Now().Add(-time.Hour).Unix(), // Token expired one hour ago
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	assert.NoError(t, err)

	claims, err := jwtService.ValidateToken(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
}
