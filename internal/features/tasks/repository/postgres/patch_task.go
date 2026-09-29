package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *TasksRepository) PatchTask(
	ctx context.Context,
	id int64,
	userID int64,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.task
	SET
		title=@title,
		description=@description,
		completed=@completed,
		completed_at=@completed_at,
		version=version+1
	WHERE id=@id AND user_id=@user_id AND version=@version;`
	args := pgx.NamedArgs{
		"title":        task.Title,
		"description":  task.Description,
		"completed":    task.Completed,
		"completed_at": task.CompletedAt,
	}
	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return domain.Task{}, fmt.Errorf("query task: %w", err)
	}

	model, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[taskModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d' concurrently accessed: %w",
				core_errors.ErrConflict,
			)
		}
		return domain.Task{}, fmt.Errorf("collect task: %w", err)
	}

	return model.ToDomain(), nil
}
