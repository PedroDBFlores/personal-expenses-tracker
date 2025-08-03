package main

import (
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"pedro/personal-expenses-tracker/routes"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app := fiber.New()
	routes.Setup(app)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("API running", zap.String("addr", ":"+port))
	if err := app.Listen(":" + port); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
