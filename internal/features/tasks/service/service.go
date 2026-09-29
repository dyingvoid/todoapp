package tasks_service

import (
	"context"

	"github.com/dyingvoid/todoapp/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

// PatchTask implements [tasks_transport_http.TasksService].
func (s *TasksService) PatchTask(ctx context.Context, id int64, userID int64) (domain.Task, error) {
	panic("unimplemented")
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

	PatchTask(
		ctx context.Context,
		id int64,
		userID int64,
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
