package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id int64,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.user
	WHERE id = @id;`
	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user query: %w", err)
	}

	userModel, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[UserModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("get user collect: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
