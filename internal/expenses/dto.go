package expenses

import "time"

type ExpenseRequest struct {
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Amount      float32    `json:"amount"`
	Date        *time.Time `json:"date"`
	GroupID     *string    `json:"groupId"`
}

type ExpenseListResponse struct {
	Expenses []ExpenseResponse `json:"expenses"`
}

type ExpenseResponse struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Amount      float32    `json:"amount"`
	Date        *time.Time `json:"date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	Category    string     `json:"category"`
}
