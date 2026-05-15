package groups

import (
	"time"

	"github.com/edsontoledo-g/group-expense-api/internal/domain/expenses"
)

type Group struct {
	ID          string
	Name        string
	Description string
	Expenses    *[]expenses.Expense
	CreatedAt   time.Time
}

type UserGroups struct {
	UserID  string
	GroupID string
}
