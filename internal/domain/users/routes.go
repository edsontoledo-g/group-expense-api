package users

import (
	"os"

	"github.com/edsontoledo-g/group-expense-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoutes(rg *gin.RouterGroup, h *UserHandler) {
	user := rg.Group("/me")
	user.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
	{
		user.GET("", h.GetInformation)
		user.PATCH("/groups/:id/archive", h.ArchiveGroup)
	}
}
