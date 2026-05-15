package users

import "github.com/gin-gonic/gin"

type Module struct {
	handler *UserHandler
}

func (module *Module) RegisterRoutes(rg *gin.RouterGroup) {
	registerRoutes(rg, module.handler)
}

func NewModule() *Module {
	handler := NewUserHandler()
	return &Module{
		handler: handler,
	}
}
