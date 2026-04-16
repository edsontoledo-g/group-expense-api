package routes

import (
	"github.com/edsontoledo-g/group-expense-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, userHandler *handlers.UserHandler) {
	user := rg.Group("/me")
	{
		user.GET("", userHandler.GetInformation)
		user.PATCH("/groups/:id/archive", userHandler.ArchiveGroup)
	}
}
