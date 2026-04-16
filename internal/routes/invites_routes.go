package routes

import (
	"github.com/edsontoledo-g/group-expense-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func registerInvitesRoutes(rg *gin.RouterGroup, invitesHandler *handlers.InvitesHandler) {
	invites := rg.Group("/invites")
	{
		invites.POST("/:token", invitesHandler.AcceptInvite)
	}
}
