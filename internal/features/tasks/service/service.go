package tasks_service

import (
	"context"

	"github.com/dyingvoid/todoapp/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID *int64,
		offset *int,
		limit *int,
	) ([]domain.Task, error)

	GetTask(
		ctx context.Context,
		id int64,
	) (domain.Task, error)

	DeleteTask(
		ctx context.Context,
		id int64,
	) error

	PatchTask(
		ctx context.Context,
		id int64,
		task domain.Task,
	) (domain.Task, error)
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
