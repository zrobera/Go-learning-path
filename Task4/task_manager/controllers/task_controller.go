package controllers

import (
    "net/http"
    "task_manager/data"
    "task_manager/models"
    "github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    tasks := data.GetTasks()
    c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    task, found := data.GetTaskByID(id)
    if !found {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found!"})
        return
    }
    c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    var newTask models.Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    createdTask := data.CreateTask(newTask)
    c.IndentedJSON(http.StatusCreated, createdTask)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    task, found := data.UpdateTask(id, updatedTask)
    if !found {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found!"})
        return
    }
    c.IndentedJSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    if !data.DeleteTask(id) {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found!"})
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted!"})
}
