package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGetExpensesPortForByID struct{ mock.Mock }

func (m *mockGetExpensesPortForByID) GetAllExpenses() ([]models.Expense, error) {
	args := m.Called()
	return args.Get(0).([]models.Expense), args.Error(1)
}

func (m *mockGetExpensesPortForByID) GetExpenseByID(id uint) (*models.Expense, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Expense), args.Error(1)
}

func TestGetExpenseByIDHandler_Success(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockGetExpensesPortForByID)
	handler := handlers.NewGetExpenseByIDHandler(mockPort)
	app.Get("/expenses/:id", handler.Handle)

	expense := &models.Expense{ID: 1, Amount: 10.5, Description: "Lunch", ExpenseType: "Food"}
	mockPort.On("GetExpenseByID", uint(1)).Return(expense, nil)

	req := httptest.NewRequest(http.MethodGet, "/expenses/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got models.Expense
	_ = json.NewDecoder(resp.Body).Decode(&got)
	assert.Equal(t, expense.ID, got.ID)
	mockPort.AssertExpectations(t)
}

func TestGetExpenseByIDHandler_InvalidID(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockGetExpensesPortForByID)
	handler := handlers.NewGetExpenseByIDHandler(mockPort)
	app.Get("/expenses/:id", handler.Handle)

	req := httptest.NewRequest(http.MethodGet, "/expenses/invalid", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
