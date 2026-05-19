package auth

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrInvalidCredentials       = errors.New("invalid email or password")
	ErrAccountNotVerified       = errors.New("account is not verified")
	ErrUserNotFound             = errors.New("user not found")
	ErrInvalidVerificationToken = errors.New("verification token is invalid or has expired")
)
