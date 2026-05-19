package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	handler *GroupsHandler
}

func (module *Module) RegisterRoutes(rg *gin.RouterGroup) {
	registerRoutes(rg, module.handler)
}

func NewModule(pool *pgxpool.Pool) *Module {
	repo := NewGroupsRepository(pool)
	service := NewGroupsService(repo)
	handler := NewGroupsHandler(service)
	return &Module{
		handler: handler,
	}
}
