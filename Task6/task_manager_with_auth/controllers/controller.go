package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid input data"})
		return
	}

	err := data.CreateUser(user)
	if err != nil{
		if err.Error() == "password length must be greater than 4" {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	token, err := data.Login(user)
	if err != nil {
		if err.Error() == "password length must be greater than 4" {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(401, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User logged in successfully", "token": token})
}

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		if err.Error() ==  "mongo: no documents in result" {
			c.IndentedJSON(404, gin.H{"error": "Task not found"})	
		}else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	createdTask, err := data.CreateTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdTask)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.IndentedJSON(404, gin.H{"error": "Task not found"})	
		}else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	
	c.IndentedJSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func PromoteUser(c *gin.Context) {
	username := c.Param("username")

	_, err := data.PromoteUser(username)
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
