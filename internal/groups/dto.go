package groups

import (
	"time"

	"github.com/edsontoledo-g/group-expense-api/internal/expenses"
	"github.com/edsontoledo-g/group-expense-api/internal/users"
)

type GroupRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	UserIDs     []string `json:"userIds"`
}

type GroupInviteRequest struct {
	Token string `json:"token"`
}

type GroupResponseList struct {
	Groups []GroupResponse `json:"groups"`
}

type GroupResponse struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description *string                     `json:"description,omitempty"`
	Users       []users.UserResponse        `json:"users"`
	Expenses    *[]expenses.ExpenseResponse `json:"expenses,omitempty"`
	CreatedAt   time.Time
}
