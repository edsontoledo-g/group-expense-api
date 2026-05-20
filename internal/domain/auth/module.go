package auth

import (
	"embed"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed templates
var templateFS embed.FS

type Module struct {
	handler *AuthHandler
}

func (module *Module) RegisterRoutes(rg *gin.RouterGroup) {
	registerRoutes(rg, module.handler)
}

func Templates() fs.FS {
	sub, err := fs.Sub(templateFS, "templates")
	if err != nil {
		panic(err)
	}
	return sub
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
