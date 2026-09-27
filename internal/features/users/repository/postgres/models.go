package users_postgres_repository

import "github.com/dyingvoid/todoapp/internal/core/domain"

type UserModel struct {
	ID          int64   `db:"id"`
	Version     int64   `db:"version"`
	FullName    string  `db:"full_name"`
	PhoneNumber *string `db:"phone_number"`
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}
	return userDomains
}
