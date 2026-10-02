package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.user (full_name, phone_number)
	VALUES (@full_name, @phone_number)
	RETURNING *;`
	args := pgx.NamedArgs{
		"full_name":    user.FullName,
		"phone_number": user.PhoneNumber,
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user query: %w", err)
	}

	userModel, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[UserModel])
	if err != nil {
		return domain.User{}, fmt.Errorf("create user collect: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
