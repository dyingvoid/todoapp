package tasks_transport_http

import (
	"time"

	"github.com/dyingvoid/todoapp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int64      `json:"id"`
	Version      int64      `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int64      `json:"author_user_id"`
}

func taskFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksFromDomains(tasks []domain.Task) []TaskDTOResponse {
	dtos := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		dtos[i] = taskFromDomain(task)
	}
	return dtos
}
