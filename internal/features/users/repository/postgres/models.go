package users_postgres_repository

type UserModel struct {
	ID          int64   `db:"id"`
	Version     int64   `db:"version"`
	FullName    string  `db:"full_name"`
	PhoneNumber *string `db:"phone_number"`
}
