package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	id int64,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT *
	FROM todoapp.task
	WHERE id = @id;`
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return domain.Task{}, fmt.Errorf("task query: %w", err)
	}

	model, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[taskModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task not found: %v: %w",
				err,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("collect task: %w", err)
	}

	return model.ToDomain(), nil
}
