package models

type Group struct {
	ID          string
	Name        string
	Description string
	UserIDs     []string
	Expenses    *[]Expense
}
