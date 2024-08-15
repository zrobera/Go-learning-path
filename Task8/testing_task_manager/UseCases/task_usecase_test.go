package usecases_test

import (
	"context"
	"errors"
	domain "test_task_manager/Domain"
	usecases "test_task_manager/UseCases"
	mocks "test_task_manager/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskUseCaseSuite struct {
	suite.Suite
	taskRepository *mocks.TaskRepository
	taskUseCase    domain.TaskUseCase
}

func (suite *TaskUseCaseSuite) SetupTest() {
	// Create a new mock TaskRepository
	suite.taskRepository = new(mocks.TaskRepository)

	suite.taskUseCase = usecases.NewTaskUseCase(suite.taskRepository, 2*time.Second)
}

func (suite *TaskUseCaseSuite) TestCreateTask_Positive() {
	task := domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	suite.taskRepository.On("CreateTask", mock.Anything, task).Return(&task, nil)

	createdTask, err := suite.taskUseCase.CreateTask(context.Background(), task)

	suite.NoError(err)
	suite.NotNil(createdTask)
	suite.Equal(task.ID, createdTask.ID)
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestCreateTask_Negative() {
	task := domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	suite.taskRepository.On("CreateTask", mock.Anything, task).Return(nil, errors.New("failed to create task"))

	createdTask, err := suite.taskUseCase.CreateTask(context.Background(), task)

	suite.Error(err)
	suite.Nil(createdTask)
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestDeleteTask_Positive() {
	taskID := "1"

	suite.taskRepository.On("DeleteTask", mock.Anything, taskID).Return(nil)

	err := suite.taskUseCase.DeleteTask(context.Background(), taskID)

	suite.NoError(err)                              
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestDeleteTask_Negative() {
	taskID := "1"

	suite.taskRepository.On("DeleteTask", mock.Anything, taskID).Return(errors.New("failed to delete task"))

	err := suite.taskUseCase.DeleteTask(context.Background(), taskID)

	suite.Error(err)                          
	suite.taskRepository.AssertExpectations(suite.T())
}


func (suite *TaskUseCaseSuite) TestGetTaskByID_Positive() {
	taskID := "1"
	task := domain.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	suite.taskRepository.On("GetTaskByID", mock.Anything, taskID).Return(&task, nil)

	fetchedTask, err := suite.taskUseCase.GetTaskByID(context.Background(), taskID)

	suite.NoError(err)
	suite.NotNil(fetchedTask)
	suite.Equal(task.ID, fetchedTask.ID)
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetTaskByID_Negative() {
	taskID := "1"

	suite.taskRepository.On("GetTaskByID", mock.Anything, taskID).Return(nil, errors.New("task not found"))

	fetchedTask, err := suite.taskUseCase.GetTaskByID(context.Background(), taskID)

	suite.Error(err)
	suite.Nil(fetchedTask)
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetTasks_Positive() {
	tasks := []domain.Task{
		{
			ID:          "1",
			Title:       "Test Task 1",
			Description: "This is a test task 1",
			Status:      "pending",
			DueDate:     time.Now().Add(24 * time.Hour),
		},
		{
			ID:          "2",
			Title:       "Test Task 2",
			Description: "This is a test task 2",
			Status:      "completed",
			DueDate:     time.Now().Add(48 * time.Hour),
		},
	}

	suite.taskRepository.On("GetTasks", mock.Anything).Return(tasks, nil)

	retrievedTasks, err := suite.taskUseCase.GetTasks(context.Background())

	suite.NoError(err)                                
	suite.NotNil(retrievedTasks)     
	suite.Equal(len(tasks), len(retrievedTasks))
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetTasks_Negative() {

	suite.taskRepository.On("GetTasks", mock.Anything).Return(nil, errors.New("failed to retrieve tasks"))

	retrievedTasks, err := suite.taskUseCase.GetTasks(context.Background())

	suite.Error(err)                           
	suite.Nil(retrievedTasks)                        
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestUpdateTask_Positive() {
	taskID := "1"
	updatedTask := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "This is an updated test task",
		Status:      "completed",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	suite.taskRepository.On("UpdateTask", mock.Anything, taskID, updatedTask).Return(&updatedTask, nil)

	result, err := suite.taskUseCase.UpdateTask(context.Background(), taskID, updatedTask)

	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(updatedTask.ID, result.ID)
	suite.Equal(updatedTask.Title, result.Title)
	suite.taskRepository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestUpdateTask_Negative() {
	taskID := "1"
	updatedTask := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "This is an updated test task",
		Status:      "completed",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	suite.taskRepository.On("UpdateTask", mock.Anything, taskID, updatedTask).Return(nil, errors.New("failed to update task"))

	result, err := suite.taskUseCase.UpdateTask(context.Background(), taskID, updatedTask)

	suite.Error(err)                                
	suite.Nil(result)
	suite.taskRepository.AssertExpectations(suite.T())
}

func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseSuite))
}
