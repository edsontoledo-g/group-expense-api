package auth

import "time"

type AuthProvider struct {
	ID                         uint
	UserID                     uint
	Provider                   string
	ProviderUserID             *string
	PasswordHash               *string
	VerificationTokenHash      *string
	VerificationTokenExpiresAt *time.Time
	CreatedAt                  time.Time
}
