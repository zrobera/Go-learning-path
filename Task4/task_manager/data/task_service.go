package data

import (
	"task_manager/models"
	"time"
)

// In-memory database
var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (*models.Task, bool) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, true
		}
	}
	return nil, false
}

func CreateTask(newTask models.Task) models.Task {
	tasks = append(tasks, newTask)
	return newTask
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, bool) {
	for i, task := range tasks {
		if task.ID == id {

			// update only the fields that are provided through the updatedTask data
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			return &tasks[i], true
		}
	}
	return nil, false
}

func DeleteTask(id string) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
