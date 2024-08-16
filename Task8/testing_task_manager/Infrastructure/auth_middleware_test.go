package infrastructure_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	infrastructure "test_task_manager/Infrastructure"
	mocks "test_task_manager/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareTestSuite struct {
	suite.Suite
	jwtService     *mocks.JWTService
	authMiddleware *infrastructure.AuthMiddleware
	router         *gin.Engine
}

func (suite *AuthMiddlewareTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET", "testsecret")
}

func (suite *AuthMiddlewareTestSuite) SetupTest() {
	suite.jwtService = new(mocks.JWTService)
	suite.authMiddleware = infrastructure.NewAuthMiddleware(suite.jwtService)
}

func (suite *AuthMiddlewareTestSuite) setupRouter(adminOnly bool, route string) {
	suite.router = gin.New()
	suite.router.Use(suite.authMiddleware.AuthMiddleware(adminOnly))
	suite.router.GET(route, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
}

func (suite *AuthMiddlewareTestSuite) TestNoAuthorizationHeader() {
	suite.setupRouter(false, "/test")

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(suite.T(), `{"error":"Authorization header is required"}`, w.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestMalformedAuthorizationHeader() {
	suite.setupRouter(false, "/test")

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "InvalidToken")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(suite.T(), `{"error":"Invalid authorization header"}`, w.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestExpiredToken() {
	suite.setupRouter(false, "/test")

	suite.jwtService.On("ValidateToken", mock.Anything).Return(nil, errors.New("Token is expired"))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer expiredToken")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(suite.T(), `{"error":"Token is expired"}`, w.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestUserRoleAccessNonAdminEndpoint() {
	suite.setupRouter(false, "/test")

	suite.jwtService.On("ValidateToken", mock.Anything).Return(map[string]interface{}{"username": "testuser", "role": "User"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer validTokenForUser")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"success"}`, w.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestUserRoleAccessAdminOnlyEndpoint() {
	suite.setupRouter(true, "/admin")

	suite.jwtService.On("ValidateToken", mock.Anything).Return(map[string]interface{}{"username": "testuser", "role": "User"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer validTokenForUser")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusForbidden, w.Code)
	assert.JSONEq(suite.T(), `{"error":"User role not allowed to access this endpoint"}`, w.Body.String())
}

func (suite *AuthMiddlewareTestSuite) TestAdminRoleAccessAdminOnlyEndpoint() {
	suite.setupRouter(true, "/admin")

	suite.jwtService.On("ValidateToken", mock.Anything).Return(map[string]interface{}{"username": "testuser", "role": "Admin"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer validTokenForAdmin")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"success"}`, w.Body.String())
}

func TestAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
