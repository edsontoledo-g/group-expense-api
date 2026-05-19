package expenses

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	handler *ExpensesHandler
}

func (module *Module) RegisterRoutes(rg *gin.RouterGroup) {
	registerRoutes(rg, module.handler)
}

func NewModule(db *pgxpool.Pool) *Module {
	repo := NewExpensesRepository(db)
	service := NewExpensesService(repo)
	handler := NewExpensesHandler(service)
	return &Module{
		handler: handler,
	}
}
