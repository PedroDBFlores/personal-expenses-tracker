package usecases

import (
	"gorm.io/gorm"
)

type DeleteExpenseUseCase struct {
	DB *gorm.DB
}

func NewDeleteExpenseUseCase(db *gorm.DB) *DeleteExpenseUseCase {
	return &DeleteExpenseUseCase{DB: db}
}

func (uc *DeleteExpenseUseCase) DeleteExpense(id uint) error {
	return uc.DB.Delete(&struct{ ID uint }{ID: id}).Error
}
