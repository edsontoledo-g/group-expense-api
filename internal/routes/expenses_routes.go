package routes

import (
	"github.com/edsontoledo-g/group-expense-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func registerExpensesRoutes(rg *gin.RouterGroup, expensesHandler *handlers.ExpensesHandler) {
	expenses := rg.Group("expenses")
	{
		expenses.GET("", expensesHandler.GetExpensesHandler)
		expenses.POST("", expensesHandler.CreateExpenseHandler)
		expenses.DELETE("/:id", expensesHandler.DeleteExpenseHandler)
	}
}
