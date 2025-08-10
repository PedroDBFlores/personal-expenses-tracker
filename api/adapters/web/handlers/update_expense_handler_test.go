package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUpdateExpensePortForUpdate struct{ mock.Mock }

func (m *mockUpdateExpensePortForUpdate) UpdateExpense(id uint, req payloads.UpdateExpenseRequest) (*models.Expense, error) {
	args := m.Called(id, req)
	return args.Get(0).(*models.Expense), args.Error(1)
}

func TestUpdateExpenseHandler_Success(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockUpdateExpensePortForUpdate)
	handler := handlers.NewUpdateExpenseHandler(mockPort)
	app.Put("/expenses/:id", handler.Handle)

	reqBody := `{"amount":15.5}`
	updatedExpense := &models.Expense{ID: 1, Amount: 15.5, Description: "Updated Lunch", ExpenseType: "Food"}
	mockPort.On("UpdateExpense", uint(1), mock.MatchedBy(func(req payloads.UpdateExpenseRequest) bool {
		return req.Amount != nil && *req.Amount == 15.5
	})).Return(updatedExpense, nil)

	req := httptest.NewRequest(http.MethodPut, "/expenses/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got models.Expense
	_ = json.NewDecoder(resp.Body).Decode(&got)
	assert.Equal(t, updatedExpense.Amount, got.Amount)
	mockPort.AssertExpectations(t)
}

func TestUpdateExpenseHandler_InvalidID(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockUpdateExpensePortForUpdate)
	handler := handlers.NewUpdateExpenseHandler(mockPort)
	app.Put("/expenses/:id", handler.Handle)

	reqBody := `{"amount":15.5}`
	req := httptest.NewRequest(http.MethodPut, "/expenses/invalid", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
