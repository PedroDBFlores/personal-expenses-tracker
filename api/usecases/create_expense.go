package usecases

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"

	"gorm.io/gorm"
)

type CreateExpenseUseCase struct {
	DB *gorm.DB
}

func NewCreateExpenseUseCase(db *gorm.DB) *CreateExpenseUseCase {
	return &CreateExpenseUseCase{DB: db}
}

func (uc *CreateExpenseUseCase) CreateExpense(req payloads.CreateExpenseRequest) (*models.Expense, error) {
	exp := &models.Expense{
		Amount:            req.Amount,
		Description:       req.Description,
		Date:              req.Date, // now time.Time
		ExpenseType:       req.ExpenseType,
		FulfillsExpenseId: req.FulfillsExpenseId,
	}
	if err := uc.DB.Create(exp).Error; err != nil {
		return nil, err
	}
	return exp, nil
}
