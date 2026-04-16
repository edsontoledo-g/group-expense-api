package groups

import "github.com/edsontoledo-g/group-expense-api/internal/expenses"

type Group struct {
	ID          string
	Name        string
	Description string
	UserIDs     []string
	Expenses    *[]expenses.Expense
}
