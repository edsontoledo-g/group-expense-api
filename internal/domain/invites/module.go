package invites

import "github.com/gin-gonic/gin"

type Module struct {
	handler *InvitesHandler
}

func (module *Module) RegisterRoutes(rg *gin.RouterGroup) {
	registerRoutes(rg, module.handler)
}

func NewModule() *Module {
	handler := NewInvitesHandler()
	return &Module{
		handler: handler,
	}
}
