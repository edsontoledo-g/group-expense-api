package expenses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpensesHandler struct{}

func (_ *ExpensesHandler) GetExpensesHandler(c *gin.Context) {
	// TODO: Implement
	expenseList := ExpenseListResponse{
		Expenses: []ExpenseResponse{
			{
				Id:          "xxx",
				Name:        "Test",
				Description: nil,
				Amount:      129.5,
				Date:        "29/03/1999",
				Category:    "Food",
			},
		},
	}
	c.JSON(http.StatusOK, expenseList)
}

func (_ *ExpensesHandler) CreateExpenseHandler(c *gin.Context) {
	var expense ExpenseRequest
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	// TODO: Implement
	c.JSON(http.StatusOK, expense)
}

func (_ *ExpensesHandler) DeleteExpenseHandler(c *gin.Context) {

}
