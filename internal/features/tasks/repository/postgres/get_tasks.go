package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	userID int64,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT *
	FROM todoapp.task
	WHERE user_id = @user_id
	ORDER BY id ASC
	LIMIT @limit
	OFFSET @offset;`
	args := pgx.NamedArgs{
		"user_id": userID,
		"limit":   limit,
		"offset":  offset,
	}
	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("tasks query: %w", err)
	}

	tasksModels, err := pgx.CollectRows(rows, pgx.RowToStructByName[taskModel])
	if err != nil {
		return nil, fmt.Errorf("tasks collect: %w", err)
	}

	return modelsToDomains(tasksModels), nil
}
