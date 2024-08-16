package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"test_task_manager/Delivery/controllers"
	domain "test_task_manager/Domain"
	mocks "test_task_manager/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	userUseCase *mocks.UserUseCase
	router      *gin.Engine
	controller  *controllers.UserController
}

func (suite *UserControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.userUseCase = new(mocks.UserUseCase)
	suite.controller = &controllers.UserController{UserUseCase: suite.userUseCase}
	suite.router = gin.New()
	suite.router.POST("/register", suite.controller.Register)
	suite.router.POST("/login", suite.controller.Login)
	suite.router.PUT("/promote/:username", suite.controller.PromoteUser)
}

func (suite *UserControllerTestSuite) TestRegisterPositive() {
	
	user := domain.User{Username: "testuser", Password: "password"}

	userJSON, err := json.Marshal(user)
	if err != nil {
		suite.T().Fatalf("Failed to marshal user data: %v", err)
	}

	suite.userUseCase.On("CreateUser", mock.Anything, user).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message": "user registered successfully"}`, w.Body.String())
}

func (suite *UserControllerTestSuite) TestRegisterNegative() {
	user := domain.User{Username: "testuser", Password: "short", Role: "User"}
	suite.userUseCase.On("CreateUser", mock.Anything, user).Return(errors.New("password length must be greater than 4"))

	userJSON, err := json.Marshal(user)
	if err != nil {
		suite.T().Fatalf("Failed to marshal user data: %v", err)
	}

	suite.userUseCase.On("CreateUser", mock.Anything, user).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error": "password length must be greater than 4"}`, w.Body.String())
}

func (suite *UserControllerTestSuite) TestLoginPositive() {
	user := domain.User{Username: "testuser", Password: "short"}
	suite.userUseCase.On("Login", mock.Anything, user).Return("validToken", nil)

	userJSON, err := json.Marshal(user)
	if err != nil {
		suite.T().Fatalf("Failed to marshal user data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJSON))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message": "User logged in successfully", "token": "validToken"}`, w.Body.String())
}

func (suite *UserControllerTestSuite) TestLoginNegative() {
	user := domain.User{Username: "testuser", Password: "password"}
	suite.userUseCase.On("Login", mock.Anything, user).Return("", errors.New("invalid credentials"))

	userJSON, err := json.Marshal(user)
	if err != nil {
		suite.T().Fatalf("Failed to marshal user data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJSON))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(suite.T(), `{"error": "invalid credentials"}`, w.Body.String())
}

func (suite *UserControllerTestSuite) TestPromoteUserPositive() {
	username := "testuser"
	suite.userUseCase.On("PromoteUser", mock.Anything, username).Return(&domain.User{Username: username, Role: "Admin"}, nil)

	req := httptest.NewRequest(http.MethodPut, "/promote/testuser", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message": "User promoted successfully"}`, w.Body.String())
}

func (suite *UserControllerTestSuite) TestPromoteUserNegative() {
	username := "testuser"
	suite.userUseCase.On("PromoteUser", mock.Anything, username).Return(nil, errors.New("user not found"))

	req := httptest.NewRequest(http.MethodPut, "/promote/testuser", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
	assert.JSONEq(suite.T(), `{"error": "User not found"}`, w.Body.String())
}

type TaskControllerTestSuite struct {
	suite.Suite
	taskUseCase *mocks.TaskUseCase
	router      *gin.Engine
	controller  *controllers.TaskController
}

func (suite *TaskControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.taskUseCase = new(mocks.TaskUseCase)
	suite.controller = &controllers.TaskController{TaskUseCase: suite.taskUseCase}
	suite.router = gin.New()
	suite.router.GET("/tasks", suite.controller.GetTasks)
	suite.router.GET("/tasks/:id", suite.controller.GetTaskByID)
	suite.router.POST("/tasks", suite.controller.CreateTask)
	suite.router.PUT("/tasks/:id", suite.controller.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.controller.DeleteTask)
}

func (suite *TaskControllerTestSuite) TestGetTasksPositive() {
	tasks := []domain.Task{
		{
			ID:          "1",
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now().Round(time.Millisecond),
			Status:      "pending",
		},
	}

	suite.taskUseCase.On("GetTasks", mock.Anything).Return(tasks, nil)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	expectedResponse := `[{"id":"1","title":"Task 1","description":"Description 1","due_date":"` + tasks[0].DueDate.Format(time.RFC3339Nano) + `","status":"pending"}]`

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestGetTasksNegative() {
	expectedError := errors.New("database connection error")
	suite.taskUseCase.On("GetTasks", mock.Anything).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.JSONEq(suite.T(), `{"error": "database connection error"}`, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestGetTaskByIDPositive() {
	task := domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now().Round(time.Millisecond),
		Status:      "pending",
	}
	suite.taskUseCase.On("GetTaskByID", mock.Anything, "1").Return(&task, nil)

	req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	expectedResponse := `{"id":"1","title":"Task 1","description":"Description 1","due_date":"` + task.DueDate.Format(time.RFC3339Nano) + `","status":"pending"}`

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestGetTaskByIDNegative() {
	taskID := "1"
	expectedError := errors.New("task not found")
	suite.taskUseCase.On("GetTaskByID", mock.Anything, taskID).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.JSONEq(suite.T(), `{"error": "task not found"}`, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestCreateTaskPositive() {
	task := domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now().Round(time.Millisecond),
		Status:      "pending",
	}

	suite.taskUseCase.On("CreateTask", mock.Anything, mock.Anything).Return(&task, nil)

	taskJSON, err := json.Marshal(task)
	if err != nil {
		suite.T().Fatalf("Failed to marshal task data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	expectedResponse := `{"id":"1","title":"Task 1","description":"Description 1","due_date":"` + task.DueDate.Format(time.RFC3339Nano) + `","status":"pending"}`

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestCreateTaskInvalidInput() {
	task := domain.Task{}
	suite.taskUseCase.On("CreateTask", mock.Anything, task).Return(nil, errors.New("Invalid input data"))

	req := httptest.NewRequest(http.MethodPost, "/tasks", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error": "Invalid input data"}`, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestUpdateTaskPositive() {
	task := domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now().Round(time.Millisecond),
		Status:      "pending",
	}

	suite.taskUseCase.On("UpdateTask", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&task, nil)

	taskJSON, err := json.Marshal(task)
	if err != nil {
		suite.T().Fatalf("Failed to marshal task data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(taskJSON))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	expectedResponse := `{"id":"1","title":"Task 1","description":"Description 1","due_date":"` + task.DueDate.Format(time.RFC3339Nano) + `","status":"pending"}`

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestUpdateTaskNegative() {
	taskID := "1"
	updatedTask := domain.Task{Title: "Updated Title"}
	expectedError := errors.New("task not found")
	suite.taskUseCase.On("UpdateTask", mock.Anything, taskID, updatedTask).Return(nil, expectedError)

	taskJSON, err := json.Marshal(updatedTask)
	if err != nil {
		suite.T().Fatalf("Failed to marshal task data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(taskJSON))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.JSONEq(suite.T(), `{"error": "task not found"}`, w.Body.String())
}

func (suite *TaskControllerTestSuite) TestDeleteTaskPositive() {
	suite.taskUseCase.On("DeleteTask", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
}

func (suite *TaskControllerTestSuite) TestDeleteTaskNegative() {
	taskID := "1"
	expectedError := errors.New("task deletion failed")
	suite.taskUseCase.On("DeleteTask", mock.Anything, taskID).Return(expectedError)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.JSONEq(suite.T(), `{"error": "task deletion failed"}`, w.Body.String())
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func TestTaskController(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
