package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	domain "test_task_manager/Domain"
	usecases "test_task_manager/UseCases"
	mocks "test_task_manager/mocks"
)

type UserUseCaseSuite struct {
	suite.Suite
	userRepository  *mocks.UserRepository
	passwordService *mocks.PasswordService
	jwtService      *mocks.JWTService
	userUseCase     domain.UserUseCase
}

func (suite *UserUseCaseSuite) SetupTest() {
	suite.userRepository = new(mocks.UserRepository)
	suite.passwordService = new(mocks.PasswordService)
	suite.jwtService = new(mocks.JWTService)

	suite.userUseCase = usecases.NewUserUseCase(suite.userRepository, suite.passwordService, suite.jwtService, 2*time.Second)
}

func (suite *UserUseCaseSuite) TestGetUsers() {
	users := []domain.User{
		{Username: "user1"},
		{Username: "user2"},
	}

	suite.userRepository.On("GetUsers", mock.Anything).Return(users, nil)

	result, err := suite.userUseCase.GetUsers(context.Background())

	suite.NoError(err)
	suite.Equal(users, result)
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUseCaseSuite) TestCreateUser_Positive() {
	user := domain.User{
		Username: "user1",
		Password: "password123",
	}

	hashedPassword := "hashedpassword"
	suite.passwordService.On("Hash", user.Password).Return(hashedPassword, nil)
	suite.userRepository.On("GetUsers", mock.Anything).Return([]domain.User{}, nil)
	suite.userRepository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{}, errors.New("user not found"))

	// Update the expected user object with the hashed password and role
	expectedUser := domain.User{
		Username: "user1",
		Password: hashedPassword,
		Role:     "Admin",
	}
	suite.userRepository.On("CreateUser", mock.Anything, expectedUser).Return(nil)

	err := suite.userUseCase.CreateUser(context.Background(), user)

	suite.NoError(err)
	suite.userRepository.AssertExpectations(suite.T())
	suite.passwordService.AssertExpectations(suite.T())
}

func (suite *UserUseCaseSuite) TestCreateUser_Negative_PasswordLength() {
	user := domain.User{
		Username: "user1",
		Password: "pwd",
	}

	err := suite.userUseCase.CreateUser(context.Background(), user)

	suite.Error(err)
	suite.EqualError(err, "password length must be greater than 4")
}

func (suite *UserUseCaseSuite) TestCreateUser_Negative_UsernameExists() {
	user := domain.User{
		Username: "user1",
		Password: "password123",
	}

	hashedPassword := "hashedpassword"
	suite.passwordService.On("Hash", user.Password).Return(hashedPassword, nil)
	suite.userRepository.On("GetUsers", mock.Anything).Return([]domain.User{}, nil)
	suite.userRepository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{Username: "user1"}, nil)

	err := suite.userUseCase.CreateUser(context.Background(), user)

	suite.Error(err)
	suite.EqualError(err, "username already exists")
}

func (suite *UserUseCaseSuite) TestLogin_Positive() {
	user := domain.User{
		Username: "user1",
		Password: "password123",
	}

	hashedPassword := "hashedpassword"
	token := "validtoken"

	suite.passwordService.On("CompareHashAndPassword", hashedPassword, user.Password).Return(nil)
	suite.userRepository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{Username: "user1", Password: hashedPassword, Role: "User"}, nil)
	suite.jwtService.On("GenerateToken", user.Username, "User").Return(token, nil)

	resultToken, err := suite.userUseCase.Login(context.Background(), user)

	suite.NoError(err)
	suite.Equal(token, resultToken)
	suite.userRepository.AssertExpectations(suite.T())
	suite.passwordService.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
}


func (suite *UserUseCaseSuite) TestLogin_Negative_User_Not_Found() {
	user := domain.User{
		Username: "user1",
		Password: "password123",
	}

	suite.userRepository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{}, errors.New("user not found"))

	resultToken, err := suite.userUseCase.Login(context.Background(), user)

	suite.Error(err)
	suite.Empty(resultToken)
	suite.EqualError(err, "invalid credentials")
}

func (suite *UserUseCaseSuite) TestLogin_Negative_PasswordMismatch() {
	user := domain.User{
		Username: "user1",
		Password: "password123",
	}

	hashedPassword := "hashedpassword"
	suite.passwordService.On("CompareHashAndPassword", hashedPassword, user.Password).Return(errors.New("password mismatch"))
	suite.userRepository.On("FindByUsername", mock.Anything, user.Username).Return(&domain.User{Username: "user1", Password: hashedPassword}, nil)

	resultToken, err := suite.userUseCase.Login(context.Background(), user)

	suite.Error(err)
	suite.Empty(resultToken)
	suite.EqualError(err, "invalid credentials")
}


func (suite *UserUseCaseSuite) TestPromoteUser_Positive() {
	username := "user1"
	user := domain.User{
		Username: username,
		Role:     "User",
	}

	suite.userRepository.On("FindByUsername", mock.Anything, username).Return(&user, nil)
	suite.userRepository.On("PromoteUser", mock.Anything, username).Return(&domain.User{Username: username, Role: "Admin"}, nil)

	result, err := suite.userUseCase.PromoteUser(context.Background(), username)

	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal("Admin", result.Role)
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUseCaseSuite) TestPromoteUser_Negative_UserNotFound() {
	username := "user1"

	suite.userRepository.On("FindByUsername", mock.Anything, username).Return(&domain.User{}, errors.New("user not found"))

	result, err := suite.userUseCase.PromoteUser(context.Background(), username)

	suite.Error(err)
	suite.Nil(result)
	suite.EqualError(err, "user not found")
}

func (suite *UserUseCaseSuite) TestPromoteUser_Negative_AlreadyAdmin() {
	username := "user1"
	user := domain.User{
		Username: username,
		Role:     "Admin",
	}

	suite.userRepository.On("FindByUsername", mock.Anything, username).Return(&user, nil)

	result, err := suite.userUseCase.PromoteUser(context.Background(), username)

	suite.Error(err)
	suite.Nil(result)
	suite.EqualError(err, "user is already an admin")
}

func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseSuite))
}
