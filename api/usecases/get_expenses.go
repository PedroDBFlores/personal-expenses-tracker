package usecases

import (
	"pedro/personal-expenses-tracker/models"

	"gorm.io/gorm"
)

type GetExpensesUseCase struct {
	DB *gorm.DB
}

func NewGetExpensesUseCase(db *gorm.DB) *GetExpensesUseCase {
	return &GetExpensesUseCase{DB: db}
}

func (uc *GetExpensesUseCase) GetAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	if err := uc.DB.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (uc *GetExpensesUseCase) GetExpenseByID(id uint) (*models.Expense, error) {
	var expense models.Expense
	if err := uc.DB.First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}
