package ports

type DeleteExpensePort interface {
	DeleteExpense(id uint) error
}
