package usecases

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"

	"gorm.io/gorm"
)

type SearchExpensesUseCase struct {
	DB *gorm.DB
}

func NewSearchExpensesUseCase(db *gorm.DB) *SearchExpensesUseCase {
	return &SearchExpensesUseCase{DB: db}
}

func (uc *SearchExpensesUseCase) SearchExpenses(req payloads.SearchExpensesRequest) ([]models.Expense, error) {
	var expenses []models.Expense
	query := uc.DB.Model(&models.Expense{})
	if req.FromDate != nil {
		query = query.Where("date >= ?", req.FromDate)
	}
	if req.ToDate != nil {
		query = query.Where("date <= ?", req.ToDate)
	}
	if req.MinAmount != nil {
		query = query.Where("amount >= ?", req.MinAmount)
	}
	if req.MaxAmount != nil {
		query = query.Where("amount <= ?", req.MaxAmount)
	}
	if req.ExpenseType != nil {
		query = query.Where("expense_type = ?", req.ExpenseType)
	}
	if err := query.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}
