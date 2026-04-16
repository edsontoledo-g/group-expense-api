package expenses

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	h := &ExpensesHandler{}
	expenses := rg.Group("/expenses")
	{
		expenses.GET("", h.GetExpensesHandler)
		expenses.POST("", h.CreateExpenseHandler)
		expenses.DELETE("/:id", h.DeleteExpenseHandler)
	}
}
