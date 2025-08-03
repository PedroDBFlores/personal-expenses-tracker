package usecases

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"

	"gorm.io/gorm"
)

type UpdateExpenseUseCase struct {
	DB *gorm.DB
}

func NewUpdateExpenseUseCase(db *gorm.DB) *UpdateExpenseUseCase {
	return &UpdateExpenseUseCase{DB: db}
}

func (uc *UpdateExpenseUseCase) UpdateExpense(id uint, req payloads.UpdateExpenseRequest) (*models.Expense, error) {
	var expense models.Expense
	if err := uc.DB.First(&expense, id).Error; err != nil {
		return nil, err
	}
	if req.Amount != nil {
		expense.Amount = *req.Amount
	}
	if req.Description != nil {
		expense.Description = *req.Description
	}
	if req.Date != nil {
		expense.Date = *req.Date // now time.Time
	}
	if req.ExpenseType != nil {
		expense.ExpenseType = *req.ExpenseType
	}
	if req.FulfillsExpenseId != nil {
		expense.FulfillsExpenseId = req.FulfillsExpenseId
	}
	if err := uc.DB.Save(&expense).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}
