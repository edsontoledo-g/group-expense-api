package auth

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	handler *AuthHandler
}

func (module *Module) RegisterRoutes(rg *gin.RouterGroup) {
	registerRoutes(rg, module.handler)
}

func NewModule(db *pgxpool.Pool) *Module {
	repo := NewAuthRepository(db)
	tokens := NewTokenService()
	mailing := NewEmailService()
	baseURL := os.Getenv("APP_BASE_URL")
	service := NewAuthService(repo, tokens, mailing, baseURL)
	handler := NewAuthHandler(service)
	return &Module{
		handler: handler,
	}
}
