package expenses

import "time"

type Expense struct {
	Id          string
	Name        string
	Description *string
	Amount      float32
	UserID      string
	Date        *time.Time
	CreatedAt   time.Time
	Category    string
	GroupID     *string
}
