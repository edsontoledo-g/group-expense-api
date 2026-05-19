package expenses

import (
	"os"

	"github.com/edsontoledo-g/group-expense-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoutes(rg *gin.RouterGroup, h *ExpensesHandler) {
	expenses := rg.Group("/expenses")
	expenses.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
	{
		expenses.GET("", h.List)
		expenses.POST("", h.Create)
		expenses.DELETE("/:id", h.Delete)
	}
}
