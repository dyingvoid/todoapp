package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.user
	ORDER BY id ASC
	LIMIT @limit
	OFFSET @offset;`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	userModels, err := pgx.CollectRows(rows, pgx.RowToStructByName[UserModel])
	if err != nil {
		return nil, fmt.Errorf("collect users: %w", err)
	}

	userDomains := userDomainsFromModels(userModels)

	return userDomains, nil
}
