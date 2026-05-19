package auth

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountNotVerified = errors.New("account is not verified")
)
