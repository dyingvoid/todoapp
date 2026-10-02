package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	core_postgres_errors "github.com/dyingvoid/todoapp/internal/core/repository/postgres/errors"
	"github.com/jackc/pgx/v5"
)

func (r *TasksRepository) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
    INSERT INTO todoapp.task (user_id, title, description, completed, created_at, completed_at)
    VALUES (@user_id, @title, @description, @completed, @created_at, @completed_at)
    RETURNING *;`
	args := pgx.NamedArgs{
		"user_id":      task.AuthorUserID,
		"title":        task.Title,
		"description":  task.Description,
		"completed":    task.Completed,
		"created_at":   task.CreatedAt,
		"completed_at": task.CompletedAt,
	}
	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		if errors.Is(err, core_postgres_errors.ErrFKViolation) {
			return domain.Task{}, fmt.Errorf(
				"%v: user with id='%d' not found: %w",
				err,
				task.AuthorUserID,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("task query: %w", err)
	}

	taskModel, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[taskModel])
	if err != nil {
		return domain.Task{}, fmt.Errorf("task collect: %w", err)
	}

	return taskModel.ToDomain(), nil
}
