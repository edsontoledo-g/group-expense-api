package dto

type ExpenseRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Amount      float32 `json:"amount"`
	Date        *string `json:"date"`
	GroupID     *string `json:"groupId"`
}
