package repositories_test

import (
	"context"
	domain "test_task_manager/Domain"
	repositories "test_task_manager/Repositories"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositorySuite struct {
	suite.Suite
	repository domain.UserRepository
	database   *mongo.Database
	cleanup    func()
}

func (suite *UserRepositorySuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	suite.Require().NoError(err)

	db := client.Database("test_db")
	suite.database = db

	repository := repositories.NewUserRepository(*db, "users")
	suite.repository = repository

	// Function to cleanup the database after each test
	suite.cleanup = func() {
		db.Collection("users").Drop(context.TODO())
	}
}

func (suite *UserRepositorySuite) TearDownTest() {
	suite.cleanup()
}

func (suite *UserRepositorySuite) TestCreateUser() {

	newUser := domain.User{
		Username: "test_user",
		Password: "hashedpassword",
		Role:     "User",
	}

	err := suite.repository.CreateUser(context.TODO(), newUser)

	suite.NoError(err)

	// Verify that the user was inserted by fetching it back
	retrievedUser, err := suite.repository.FindByUsername(context.TODO(), newUser.Username)
	suite.NoError(err)
	suite.Equal(newUser.Username, retrievedUser.Username)
	suite.Equal(newUser.Role, retrievedUser.Role)
}

func (suite *UserRepositorySuite) TestFindByUsername() {

	newUser := domain.User{
		Username: "search_user",
		Password: "search_password",
		Role:     "User",
	}

	err := suite.repository.CreateUser(context.TODO(), newUser)
	suite.Require().NoError(err)

	foundUser, err := suite.repository.FindByUsername(context.TODO(), newUser.Username)

	suite.NoError(err)
	suite.Equal(newUser.Username, foundUser.Username)
	suite.Equal(newUser.Role, foundUser.Role)
}

func (suite *UserRepositorySuite) TestFindByUsernameNotFound() {
	foundUser, err := suite.repository.FindByUsername(context.TODO(), "non_existent_user")

	// Assert that an error occurred and no user was found
	suite.Error(err)
	suite.Nil(foundUser)
}

func (suite *UserRepositorySuite) TestGetUsers() {
	users := []domain.User{
		{
			Username: "user1",
			Password: "password1",
			Role:     "User",
		},
		{
			Username: "user2",
			Password: "password2",
			Role:     "User",
		},
	}
	for _, user := range users {
		err := suite.repository.CreateUser(context.TODO(), user)
		suite.Require().NoError(err)
	}

	retrievedUsers, err := suite.repository.GetUsers(context.TODO())

	suite.NoError(err)
	suite.Equal(len(users), len(retrievedUsers))
}

func (suite *UserRepositorySuite) TestPromoteUser() {

	newUser := domain.User{
		Username: "promote_user",
		Password: "promote_password",
		Role:     "User",
	}
	err := suite.repository.CreateUser(context.TODO(), newUser)
	suite.Require().NoError(err)

	updatedUser, err := suite.repository.PromoteUser(context.TODO(), newUser.Username)

	suite.NoError(err)
	suite.Equal("Admin", updatedUser.Role)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
