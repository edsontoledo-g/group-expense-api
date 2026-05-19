package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/edsontoledo-g/group-expense-api/internal/domain/users"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(input *AuthInput) error
	SignIn(input *AuthInput) (*AuthResult, error)
	SignInWithApple(input *AuthInput) (*AuthResult, error)
	VerifyUserEmail(token string) error
}

type authService struct {
	repo    AuthRepository
	tokens  TokenService
	mailing EmailService
}

type AuthInput struct {
	Email          string
	Password       string
	Provider       string
	FirstName      *string
	LastName       *string
	ProviderUserID *string
}

type AuthResult struct {
	JWT       string
	ExpiresIn uint
	UserID    uint
}

func (s *authService) SignUp(input *AuthInput) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	hashed := string(passwordHash)
	token, err := s.tokens.GenerateEmailVerification(input.Email)
	if err != nil {
		return err
	}
	user := &users.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
	auth := &AuthProvider{
		Provider:                   input.Provider,
		ProviderUserID:             input.ProviderUserID,
		PasswordHash:               &hashed,
		VerificationTokenHash:      &token.HashedToken,
		VerificationTokenExpiresAt: &token.ExpiresAt,
	}
	err = s.repo.CreateUser(user, auth)
	if err != nil {
		return err
	}
	// TODO: Refactor hardcoded URL
	verifyURL := fmt.Sprintf(
		"http://localhost:8080/api/v1/auth/verify?token=%s",
		token.Token,
	)
	err = s.mailing.SendVerificationEmail(user.Email, verifyURL)
	return err
}

func (s *authService) SignIn(input *AuthInput) (*AuthResult, error) {
	auth, err := s.repo.GetAuthByEmail(input.Email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(*auth.PasswordHash),
		[]byte(input.Password),
	)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	user, err := s.repo.GetUserByUserID(auth.UserID)
	if err != nil {
		return nil, err
	}
	if user.IsVerified == false {
		return nil, ErrAccountNotVerified
	}
	token, err := s.tokens.GenerateJWT(auth.UserID)
	if err != nil {
		return nil, err
	}
	return &AuthResult{
		JWT:    token.Token,
		UserID: auth.UserID,
	}, nil
}

func (s *authService) SignInWithApple(input *AuthInput) (*AuthResult, error) {
	return nil, nil
}

func (s *authService) VerifyUserEmail(token string) error {
	hashed := s.tokens.HashToken(token)
	user, auth, err := s.repo.GetUserByVerificationToken(hashed)
	if err != nil {
		return err
	}
	user.IsVerified = true
	user.UpdatedAt = time.Now()
	auth.VerificationTokenHash = nil
	auth.VerificationTokenExpiresAt = nil
	err = s.repo.UpdateUser(user, auth)
	return err
}

func NewAuthService(repo AuthRepository, tokens TokenService, mailing EmailService) AuthService {
	return &authService{
		repo:    repo,
		tokens:  tokens,
		mailing: mailing,
	}
}
