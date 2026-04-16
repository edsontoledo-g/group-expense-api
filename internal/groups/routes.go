package groups

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	h := &GroupsHandler{}
	groups := rg.Group("/groups")
	{
		groups.GET("", h.GetGroups)
		groups.GET("/:id", h.GetGroup)
		groups.POST("", h.CreateGroup)
		groups.POST("/:id/invite", h.CreateInvite)
	}
}
