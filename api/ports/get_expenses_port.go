package ports

import "pedro/personal-expenses-tracker/models"

type GetExpensesPort interface {
	GetAllExpenses() ([]models.Expense, error)
	GetExpenseByID(id uint) (*models.Expense, error)
}
