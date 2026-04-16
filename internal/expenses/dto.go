package expenses

type ExpenseRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Amount      float32 `json:"amount"`
	Date        *string `json:"date"`
	GroupID     *string `json:"groupId"`
}

type ExpenseListResponse struct {
	Expenses []ExpenseResponse `json:"expenses"`
}

type ExpenseResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Amount      float32 `json:"amount"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
}
