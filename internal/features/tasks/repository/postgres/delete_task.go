package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *TasksRepository) DeleteTask(
	ctx context.Context,
	id int64,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.task
	WHERE id = @id;`
	args := pgx.NamedArgs{
		"id": id,
	}

	cmdTag, err := r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"task with id='%d' not found: %w",
			id,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
