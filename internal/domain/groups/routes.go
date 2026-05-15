package groups

import (
	"os"

	"github.com/edsontoledo-g/group-expense-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoutes(rg *gin.RouterGroup, h *GroupsHandler) {
	groups := rg.Group("/groups")
	groups.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
	{
		groups.GET("", h.List)
		groups.GET("/:id", h.Get)
		groups.POST("", h.Create)
		groups.POST("/:id/invite", h.Invite)
	}
}
