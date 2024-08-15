package domain

import (
	"context"
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"` // "Admin" || "User"
}

type TaskUseCase interface {
	GetTasks(c context.Context) ([]Task, error)
	GetTaskByID(c context.Context, taskID string) (*Task, error)
	CreateTask(c context.Context, newTask Task) (*Task,error)
	UpdateTask(c context.Context, taskID string, updatedTask Task) (*Task, error)
	DeleteTask(c context.Context, taskID string) error
}

type TaskRepository interface {
	GetTasks(c context.Context) ([]Task, error)
	GetTaskByID(c context.Context, taskID string) (*Task, error)
	CreateTask(c context.Context, newTask Task) (*Task,error)
	UpdateTask(c context.Context, taskID string, updatedTask Task) (*Task, error)
	DeleteTask(c context.Context, taskID string) error
}

type UserUseCase interface {
	GetUsers(c context.Context) ([]User, error)
	CreateUser(c context.Context, user User) error
	Login(c context.Context, user User) (string, error)
	PromoteUser(c context.Context, username string) (*User, error)
}

type UserRepository interface {
	GetUsers(c context.Context) ([]User, error)
	CreateUser(c context.Context, user User) error
	FindByUsername(c context.Context, username string) (*User, error)
	PromoteUser(c context.Context, username string) (*User, error)
}