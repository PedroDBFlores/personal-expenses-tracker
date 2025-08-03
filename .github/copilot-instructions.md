# Project Overview

This project is a web application that allows a user to manage their personal expenses. It is built using React and Golang, and uses SQLite for data storage.

# Project Structure
The project is structured as follows:

```
personal-expenses-tracker/
├── api/                  # Backend API written in Golang
│   ├── go.mod            # Go module file for dependencies
│   ├── go.sum            # Go module checksum file
│   ├── main.go           # Entry point for the API server
│   ├── adapters/web/handlers      # Contains HTTP handlers for API endpoints
│   ├── adapters/web/payloads    # Payloads for request and response bodies
│   ├── adapters/db       # Database adapters and migrations
│   ├── adapters/logger   # Logging setup and configuration
│   ├── ports/            # Interfaces for the application
│   ├── usecases/         # Business logic and service layer
│   ├── models/           # Data models and database interactions                                          
│   ├── routes/           # Defines API routes and middleware
│   └── utils/            # Utility functions and helpers
├── client/               # Frontend application written in React
│   ├── src/              # Source code for the React application
│   │   ├── components/   # React components
│   │   ├── pages/        # Page components for routing
│   │   ├── services/     # API service calls, so that it's easier to test as well, and not to have tightly coupled components
│   │   └── utils/        # Utility functions and helpers
│   ├── public/           # Public assets and index.html
│   └── package.json       # NPM dependencies and scripts
└── README.md             # Project documentation
```

# Model
The application uses the following model for expenses:
- `ID`: Unique identifier for the expense (auto-incremented).
- `Amount`: The amount of the expense.
- `Description`: A brief description of the expense.
- `Date`: The timestamp of the expense.
- `CreatedAt`: Timestamp for when the expense was created (automatically set on creation).
- `UpdatedAt`: Timestamp for when the expense was last updated (automatically set on update).
- `ExpenseType`: The type of the expense, which can be either "credit" or "debit".
- `FulfillsExpenseId`: An Id of the expense that will be fulfilled by this one. This field is nullable, meaning it can be left empty if not applicable.

## Database
The application uses SQLite as the database for storing expenses. The database schema is autogenerate by GORM.

# Backend API (Golang)
The backend API is built using Golang and serves as the server-side component of the application. It lives on the `api/` directory. Golang version required is 1.24.
It uses the following libraries:
- `fiber`: A web framework for building APIs in Golang.
- `gorm`: An ORM for Golang to interact with the SQLite database.
- `sqlite3`: SQLite driver for Golang.
- `zap`: A logging library for structured logging. Instructions are available on: [https://docs.gofiber.io/contrib/fiberzap/](https://docs.gofiber.io/contrib/fiberzap/)

Currently it doesn't have any authentication or authorization implemented, but it can be added later.
- The projects import path is `pedro/personal-expenses-tracker/`, so you can import packages like this: `import "pedro/personal-expenses-tracker/models"`.

## Code structure
We're following a clean architecture approach, this allows for better separation of concerns and easier testing. An example can be found on: [https://dev.to/leapcell/clean-architecture-in-go-a-practical-guide-with-go-clean-arch-51h7](https://dev.to/leapcell/clean-architecture-in-go-a-practical-guide-with-go-clean-arch-51h7).

For example, a handler should have the following structure:

```go
package handlers
import (
    "github.com/gofiber/fiber/v2"
    "personal-expenses-tracker/api/models"
    "personal-expenses-tracker/api/utils"
)

type ExpenseHandler struct {
    service models.ExpenseService // Service for handling expense-related operations
}

func NewExpenseHandler(service models.ExpenseService) *ExpenseHandler {
    return &ExpenseHandler{
        service: service,
    }
}
func (h *ExpenseHandler) GetExpenses(c *fiber.Ctx) error {
    expenses, err := h.service.GetAllExpenses()
    if err != nil {
        return utils.HandleError(c, err)
    }
    return c.JSON(expenses)
}
```
- Ports should be defined in the `ports/` directory, which contains interfaces that define the contract for the application. The ports should be named after the use case they represent, such as `ExpenseCreator` or `ExpenseGetter`.
- Use cases should be defined in the `usecases/` directory, which contains the business logic and service layer.
- Models should be defined in the `models/` directory, which contains the data models and database interactions.
- Adapters should be defined in the `adapters/` directory, which contains the implementation of the ports and use cases.
- Routes should be defined in the `routes/` directory, which contains the API routes and middleware. All routes should be prefixed with `/api`.

Note: Each use case should have its own file in the `usecases/` directory, and the file name should match the use case name. For example, `create_expense.go` for creating an expense, `get_all_expenses.go` for getting all expenses, etc.
Note: Each handler should have its own file in the `adapters/web/` directory, and the file name should match the handler name. For example, `expense_handler.go` for handling expense-related operations.

## Endpoints
The API should have the following endpoints:
- `GET /expenses`: Retrieves all expenses.
- `POST /expenses`: Creates a new expense.
- `GET /expenses/:id`: Retrieves a specific expense by ID.
- `PUT /expenses/:id`: Updates a specific expense by ID.
- `DELETE /expenses/:id`: Deletes a specific expense by ID.
- `POST /expenses/search`: Searches for expenses based on a query. The query can include filters such as date range, amount range, and expense type.

## Testing
The backend API should be tested using the `testing` package in Golang, alongside `stretchr/testify` for assertions and mocking.
The tests should be written on a _test.go file in the same package as the code being tested.

## Deployment
The backend API can be deployed using Docker. A `Dockerfile` is provided in the `api/` directory to build the Docker image. The image can be run using the following command:
```bash
docker build -t personal-expenses-tracker-api .
docker run -p 8080:8080 personal-expenses-tracker-api
```

The image should be multistaged to reduce the final image size. The first stage should build the application, and the second stage should copy the built binary to a minimal base image.
The minimal base image should be `scratch`. The Dockerfile should also include the necessary environment variables for the application to run, such as the database connection string.

# Frontend Application (React)
The frontend application is built using React and serves as the client-side component of the application, using Typescript for type safety. It lives on the `client/` directory.
We can bootstrap it by using `vite` with the `react-ts` template.
It uses the following libraries:
- `react`: The core library for building user interfaces.
- `react-router-dom`: For routing and navigation within the application.
- `fetch`: For making HTTP requests to the backend API.
- `Material-UI`: A popular React UI framework for building responsive and modern user interfaces.

## Testing
The frontend application should be tested using `Vitest` and `React Testing Library`.
The tests should be written in a `*.test.ts` file in the same directory as the component being tested.

# Notes for Copilot
- Never suggest code changes to tests that make them pass without proper assertions.
- Ensure that when writing Go files, you only have one package name per file, and it matches the directory name.
- When writing handlers and use cases, ensure that they are properly structured and follow the clean architecture principles, writing tests for them.