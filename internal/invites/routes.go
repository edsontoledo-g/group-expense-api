package invites

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	h := &InvitesHandler{}
	invites := rg.Group("/invites")
	{
		invites.POST("/:token", h.AcceptInvite)
	}
}
