package controllers

import (
	"net/http"
	domain "task_manager/Domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUseCase domain.TaskUseCase
}

type UserController struct {
	UserUseCase domain.UserUseCase
}

// user controllers
func (u *UserController) Register(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input data"})
		return
	}

	err := u.UserUseCase.CreateUser(c, user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (u *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	token, err := u.UserUseCase.Login(c, user)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func (u *UserController) PromoteUser(c *gin.Context) {
	username := c.Param("username")

	_, err := u.UserUseCase.PromoteUser(c,username)
	if err != nil {
		if err.Error() == "user not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}


// task controllers
func (t *TaskController) GetTasks(c *gin.Context) {
	tasks, err := t.TaskUseCase.GetTasks(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := t.TaskUseCase.GetTaskByID(c, id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.IndentedJSON(404, gin.H{"error": "Task not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	createdTask, err := t.TaskUseCase.CreateTask(c,newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdTask)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	task, err := t.TaskUseCase.UpdateTask(c,id, updatedTask)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.IndentedJSON(404, gin.H{"error": "Task not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := t.TaskUseCase.DeleteTask(c,id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}