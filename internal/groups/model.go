package groups

import (
	"time"

	"github.com/edsontoledo-g/group-expense-api/internal/expenses"
)

type Group struct {
	ID          string
	Name        string
	Description string
	UserIDs     []string
	Expenses    *[]expenses.Expense
	CreatedAt   time.Time
}
