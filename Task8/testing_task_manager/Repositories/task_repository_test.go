package repositories_test

import (
	"context"
	domain "test_task_manager/Domain"
	repositories "test_task_manager/Repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositorySuite struct {
	suite.Suite
	repository domain.TaskRepository
	database   *mongo.Database
	cleanup    func()
}

func (suite *TaskRepositorySuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	suite.Require().NoError(err)

	db := client.Database("test_db")
	suite.database = db

	repository := repositories.NewTaskRepository(*db, "tasks")
	suite.repository = repository

	suite.cleanup = func() {
		db.Collection("tasks").Drop(context.TODO())
	}
}

func (suite *TaskRepositorySuite) TearDownTest() {
	suite.cleanup()
}

func (suite *TaskRepositorySuite) TestCreateTask() {
	newTask := domain.Task{
		ID:          "123",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	createdTask, err := suite.repository.CreateTask(context.TODO(), newTask)

	suite.NoError(err)
	suite.Equal(newTask.ID, createdTask.ID) //checking if the task is created successfully
	suite.Equal(newTask.Title, createdTask.Title)
}

func (suite *TaskRepositorySuite) TestCreateTask_DuplicateID() {
	newTask := domain.Task{
		ID:          "123",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	_, err := suite.repository.CreateTask(context.TODO(), newTask)
	suite.NoError(err)

	// Try to create the same task again, expecting an error
	_, err = suite.repository.CreateTask(context.TODO(), newTask)
	suite.Error(err)
	suite.EqualError(err, "task with the given id already exists")
}

func (suite *TaskRepositorySuite) TestGetTaskByID_Positive() {
	newTask := domain.Task{
		ID:          "123",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}
	_, err := suite.repository.CreateTask(context.TODO(), newTask)
	suite.NoError(err)

	task, err := suite.repository.GetTaskByID(context.TODO(), newTask.ID)
	suite.NoError(err)
	suite.Equal(newTask.ID, task.ID)
}

func (suite *TaskRepositorySuite) TestGetTaskByID_NonExistent() {
	_, err := suite.repository.GetTaskByID(context.TODO(), "non-existent-id")
	suite.Error(err)
	suite.EqualError(err, mongo.ErrNoDocuments.Error())
}

func (suite *TaskRepositorySuite) TestUpdateTask_Positive() {
	newTask := domain.Task{
		ID:          "123",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}
	_, err := suite.repository.CreateTask(context.TODO(), newTask)
	suite.NoError(err)

	updatedTask := domain.Task{
		Title: "Updated Task",
	}
	task, err := suite.repository.UpdateTask(context.TODO(), newTask.ID, updatedTask)
	suite.NoError(err)
	suite.Equal(updatedTask.Title, task.Title)
}

func (suite *TaskRepositorySuite) TestDeleteTask_Positive() {
	newTask := domain.Task{
		ID:          "123",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}
	_, err := suite.repository.CreateTask(context.TODO(), newTask)
	suite.NoError(err)

	err = suite.repository.DeleteTask(context.TODO(), newTask.ID)
	suite.NoError(err)

	// Verify that the task no longer exists
	_, err = suite.repository.GetTaskByID(context.TODO(), newTask.ID)
	suite.Error(err)
	suite.EqualError(err, mongo.ErrNoDocuments.Error())
}


func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
