package infrastructure_test

import (
	infrastructure "test_task_manager/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash_Success(t *testing.T) {
	passwordService := infrastructure.NewPasswordService()

	password := "mySecurePassword"
	hashedPassword, err := passwordService.Hash(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)
}

func TestCompareHashAndPassword_Success(t *testing.T) {
	passwordService := infrastructure.NewPasswordService()

	password := "mySecurePassword"
	hashedPassword, err := passwordService.Hash(password)
	assert.NoError(t, err)

	err = passwordService.CompareHashAndPassword(hashedPassword, password)

	assert.NoError(t, err)
}

func TestCompareHashAndPassword_InvalidPassword(t *testing.T) {
	passwordService := infrastructure.NewPasswordService()

	password := "mySecurePassword"
	hashedPassword, err := passwordService.Hash(password)
	assert.NoError(t, err)

	err = passwordService.CompareHashAndPassword(hashedPassword, "wrongPassword")

	assert.Error(t, err)
}