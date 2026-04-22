package users

import "time"

type User struct {
	ID         string
	UserName   string
	FirstName  string
	LastName   string
	Email      string
	ImageURL   *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsVerified bool
}
