package domain

type User struct {
	ID          int64
	Version     int
	Name        string
	PhoneNumber *string
}
