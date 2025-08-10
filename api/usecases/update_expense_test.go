package usecases_test

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
	"pedro/personal-expenses-tracker/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBForUpdate(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Expense{})
	assert.NoError(t, err)
	return db
}

func TestUpdateExpenseUseCase_Success(t *testing.T) {
	db := setupTestDBForUpdate(t)
	uc := usecases.NewUpdateExpenseUseCase(db)

	// Create test data
	expense := &models.Expense{Amount: 15.0, Description: "Original", ExpenseType: "Food", Date: time.Now()}
	db.Create(expense)

	// Update
	newAmount := 25.0
	newDescription := "Updated"
	req := payloads.UpdateExpenseRequest{
		Amount:      &newAmount,
		Description: &newDescription,
	}

	updated, err := uc.UpdateExpense(expense.ID, req)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, newAmount, updated.Amount)
	assert.Equal(t, newDescription, updated.Description)
}

func TestUpdateExpenseUseCase_NotFound(t *testing.T) {
	db := setupTestDBForUpdate(t)
	uc := usecases.NewUpdateExpenseUseCase(db)

	newAmount := 25.0
	req := payloads.UpdateExpenseRequest{Amount: &newAmount}

	_, err := uc.UpdateExpense(999, req)
	assert.Error(t, err)
}

func TestUpdateExpenseUseCase_PartialUpdate(t *testing.T) {
	db := setupTestDBForUpdate(t)
	uc := usecases.NewUpdateExpenseUseCase(db)

	// Create test data
	original := &models.Expense{Amount: 15.0, Description: "Original", ExpenseType: "Food", Date: time.Now()}
	db.Create(original)

	// Update only the amount
	newAmount := 25.0
	req := payloads.UpdateExpenseRequest{Amount: &newAmount}

	updated, err := uc.UpdateExpense(original.ID, req)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, newAmount, updated.Amount)
	assert.Equal(t, original.Description, updated.Description) // Should remain unchanged
	assert.Equal(t, original.ExpenseType, updated.ExpenseType) // Should remain unchanged
}
