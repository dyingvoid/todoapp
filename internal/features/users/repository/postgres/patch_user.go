package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int64,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.user
	SET 
		full_name=@full_name,
		phone_number=@phone_number,
		version=version+1
	WHERE id=@id AND version=@version
	RETURNING *`
	args := pgx.NamedArgs{
		"full_name":    user.FullName,
		"phone_number": user.PhoneNumber,
		"id":           user.ID,
		"version":      user.Version,
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return domain.User{}, fmt.Errorf("query user: %w", err)
	}

	userModel, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[UserModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' concurrently accessed: %w",
				core_errors.ErrConflict,
			)
		}
		return domain.User{}, fmt.Errorf("collect user: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)
	return userDomain, nil
}
