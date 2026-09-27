package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.user
	WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}

	cmdTag, err := r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"user with id='%d': %w",
			id,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
