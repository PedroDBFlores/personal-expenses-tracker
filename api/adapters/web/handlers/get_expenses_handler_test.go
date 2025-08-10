package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGetExpensesPortForList struct{ mock.Mock }

func (m *mockGetExpensesPortForList) GetAllExpenses() ([]models.Expense, error) {
	args := m.Called()
	return args.Get(0).([]models.Expense), args.Error(1)
}

func (m *mockGetExpensesPortForList) GetExpenseByID(id uint) (*models.Expense, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Expense), args.Error(1)
}

func TestGetExpensesHandler_Success(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockGetExpensesPortForList)
	handler := handlers.NewGetExpensesHandler(mockPort)
	app.Get("/expenses", handler.Handle)

	expenses := []models.Expense{
		{ID: 1, Amount: 10.5, Description: "Lunch", ExpenseType: "Food"},
		{ID: 2, Amount: 5.0, Description: "Coffee", ExpenseType: "Food"},
	}
	mockPort.On("GetAllExpenses").Return(expenses, nil)

	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got []models.Expense
	_ = json.NewDecoder(resp.Body).Decode(&got)
	assert.Len(t, got, 2)
	assert.Equal(t, expenses[0].ID, got[0].ID)
	mockPort.AssertExpectations(t)
}

func TestGetExpensesHandler_Error(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockGetExpensesPortForList)
	handler := handlers.NewGetExpensesHandler(mockPort)
	app.Get("/expenses", handler.Handle)

	mockPort.On("GetAllExpenses").Return([]models.Expense{}, errors.New("database error"))

	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockPort.AssertExpectations(t)
}
