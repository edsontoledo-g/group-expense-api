package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateJWT(userID uint) (*TokenResult, error)
	GenerateEmailVerification(email string) (*TokenResult, error)
	HashToken(token string) string
}

type tokenService struct{}

type TokenResult struct {
	Token       string
	HashedToken string
	ExpiresAt   time.Time
}

func (s *tokenService) GenerateJWT(userID uint) (*TokenResult, error) {
	exp := 24 * time.Hour
	claims := jwt.MapClaims{
		"sub": strconv.FormatUint(uint64(userID), 10),
		"exp": time.Now().Add(exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}
	return &TokenResult{
		Token: signed,
	}, nil
}

func (s *tokenService) GenerateEmailVerification(email string) (*TokenResult, error) {
	exp := 1 * time.Hour
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	token := base64.URLEncoding.EncodeToString(bytes)
	hashed := s.HashToken(token)
	return &TokenResult{
		Token:       token,
		HashedToken: hashed,
		ExpiresAt:   time.Now().Add(exp),
	}, nil
}

func (s *tokenService) HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func NewTokenService() TokenService {
	return &tokenService{}
}
