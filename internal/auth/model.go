package auth

import "time"

type AuthProvider struct {
	ID             string
	UserID         string
	Provider       string
	ProviderUserID *string
	PasswordHash   *string
	CreatedAt      time.Time
}
