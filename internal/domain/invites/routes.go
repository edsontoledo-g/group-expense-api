package invites

import (
	"os"

	"github.com/edsontoledo-g/group-expense-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoutes(rg *gin.RouterGroup, h *InvitesHandler) {
	invites := rg.Group("/invites")
	invites.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
	{
		invites.POST("/:token", h.AcceptInvite)
	}
}
