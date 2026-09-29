package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_http_server "github.com/dyingvoid/todoapp/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID int64,
		offset *int,
		limit *int,
	) ([]domain.Task, error)

	GetTask(
		ctx context.Context,
		id, userID int64,
	) (domain.Task, error)

	PatchTask(
		ctx context.Context,
		id int64,
		userID int64,
	) (domain.Task, error)
}

func NewTasksHTTPHandler(
	tasksService TasksService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
	}
}
