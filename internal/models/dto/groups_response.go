package dto

type GroupResponseList struct {
	Groups []GroupResponse `json:"groups"`
}

type GroupResponse struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Users       []UserResponse     `json:"users"`
	Expenses    *[]ExpenseResponse `json:"expenses,omitempty"`
}
