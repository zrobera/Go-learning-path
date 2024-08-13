package usecases

import (
	"context"
	domain "task_manager/Domain"
	"time"
)

type taskUseCase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUseCase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (t *taskUseCase) CreateTask(c context.Context, newTask domain.Task) (*domain.Task,error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.CreateTask(ctx, newTask)
}

func (t *taskUseCase) DeleteTask(c context.Context, taskID string) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.DeleteTask(ctx, taskID)
}

func (t *taskUseCase) GetTaskByID(c context.Context, taskID string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.GetTaskByID(ctx, taskID)
}

func (t *taskUseCase) GetTasks(c context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.GetTasks(ctx)
}

func (t *taskUseCase) UpdateTask(c context.Context, taskID string, updatedTask domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.UpdateTask(ctx,taskID,updatedTask)
}
