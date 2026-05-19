package expenses

type ExpensesService interface {
	Expenses() (*ExpenseListResponse, error)
	CreateExpense(req *ExpenseRequest) error
}

type expensesService struct {
	repo ExpensesRepository
}

func (s *expensesService) Expenses() (*ExpenseListResponse, error) {
	return nil, nil
}

func (s *expensesService) CreateExpense(req *ExpenseRequest) error {
	return nil
}

func NewExpensesService(repo ExpensesRepository) ExpensesService {
	return &expensesService{
		repo: repo,
	}
}
