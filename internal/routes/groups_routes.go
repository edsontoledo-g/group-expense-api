package routes

import (
	"github.com/edsontoledo-g/group-expense-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func registerGroupsRoutes(rg *gin.RouterGroup, groupsHandler *handlers.GroupsHandler) {
	groups := rg.Group("/groups")
	{
		groups.GET("", groupsHandler.GetGroups)
		groups.GET("/:id", groupsHandler.GetGroup)
		groups.POST("", groupsHandler.CreateGroup)
		groups.POST("/:id/invite", groupsHandler.CreateInvite)
	}
}
