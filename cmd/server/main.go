package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"go-subscription-service/internal/app"
	"go-subscription-service/pkg/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	logger.Init()

	container := app.InitApp()

	logger.Infof("Config loaded: port=%s, db_url=%s", container.Config.Port, container.Config.DatabaseDSN)

	router := app.SetupRouter(container)

	server := &http.Server{
		Addr:    ":" + container.Config.Port,
		Handler: router,
	}

	go func() {
		logger.Infof("Server starting on port %s", container.Config.Port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exited cleanly")
}
