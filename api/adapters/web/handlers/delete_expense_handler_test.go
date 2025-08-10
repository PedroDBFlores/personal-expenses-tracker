package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDeleteExpensePortForDelete struct{ mock.Mock }

func (m *mockDeleteExpensePortForDelete) DeleteExpense(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestDeleteExpenseHandler_Success(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockDeleteExpensePortForDelete)
	handler := handlers.NewDeleteExpenseHandler(mockPort)
	app.Delete("/expenses/:id", handler.Handle)

	mockPort.On("DeleteExpense", uint(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/expenses/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	mockPort.AssertExpectations(t)
}

func TestDeleteExpenseHandler_InvalidID(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockDeleteExpensePortForDelete)
	handler := handlers.NewDeleteExpenseHandler(mockPort)
	app.Delete("/expenses/:id", handler.Handle)

	req := httptest.NewRequest(http.MethodDelete, "/expenses/invalid", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
