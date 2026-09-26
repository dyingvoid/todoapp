package domain

type User struct {
	ID      int64
	Version int64

	FullName    string
	PhoneNumber *string
}
