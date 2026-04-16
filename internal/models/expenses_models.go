package models

type Expense struct {
	Id          string
	Name        string
	Description *string
	Amount      float32
	UserID      string
	Date        string
	Category    string
	GroupID     *string
}
