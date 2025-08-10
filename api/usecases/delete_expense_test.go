package usecases_test

import (
	"pedro/personal-expenses-tracker/models"
	"pedro/personal-expenses-tracker/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBForDelete(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Expense{})
	assert.NoError(t, err)
	return db
}

func TestDeleteExpenseUseCase_Success(t *testing.T) {
	db := setupTestDBForDelete(t)
	uc := usecases.NewDeleteExpenseUseCase(db)

	// Create test data
	expense := &models.Expense{Amount: 15.0, Description: "To Delete", ExpenseType: "Food", Date: time.Now()}
	db.Create(expense)

	err := uc.DeleteExpense(expense.ID)
	assert.NoError(t, err)

	// Verify deletion
	var count int64
	db.Model(&models.Expense{}).Where("id = ?", expense.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestDeleteExpenseUseCase_NonExistentID(t *testing.T) {
	db := setupTestDBForDelete(t)
	uc := usecases.NewDeleteExpenseUseCase(db)

	// Try to delete non-existent expense
	err := uc.DeleteExpense(999)
	assert.NoError(t, err) // GORM doesn't return error for delete with no rows affected
}

func TestDeleteExpenseUseCase_MultipleExpenses(t *testing.T) {
	db := setupTestDBForDelete(t)
	uc := usecases.NewDeleteExpenseUseCase(db)

	// Create multiple test expenses
	expense1 := &models.Expense{Amount: 15.0, Description: "Expense 1", ExpenseType: "Food", Date: time.Now()}
	expense2 := &models.Expense{Amount: 25.0, Description: "Expense 2", ExpenseType: "Transport", Date: time.Now()}
	db.Create(expense1)
	db.Create(expense2)

	// Delete one expense
	err := uc.DeleteExpense(expense1.ID)
	assert.NoError(t, err)

	// Verify only the targeted expense was deleted
	var count1 int64
	var count2 int64
	db.Model(&models.Expense{}).Where("id = ?", expense1.ID).Count(&count1)
	db.Model(&models.Expense{}).Where("id = ?", expense2.ID).Count(&count2)
	assert.Equal(t, int64(0), count1) // expense1 should be deleted
	assert.Equal(t, int64(1), count2) // expense2 should still exist
}
