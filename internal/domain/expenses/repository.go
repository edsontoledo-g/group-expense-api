package expenses

import "github.com/jackc/pgx/v5/pgxpool"

type ExpensesRepository interface {
	Create(expense *Expense) error
}

type expensesRepository struct {
	pool *pgxpool.Pool
}

func (repo *expensesRepository) Create(expense *Expense) error {
	return nil
}

func NewExpensesRepository(db *pgxpool.Pool) ExpensesRepository {
	return &expensesRepository{
		pool: db,
	}
}
