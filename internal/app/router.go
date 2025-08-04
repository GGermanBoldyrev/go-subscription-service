package app

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-subscription-service/internal/middleware"
)

func SetupRouter(container *AppContainer) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestLogger)

	// Swagger
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	container.SubscriptionHandler.RegisterRoutes(router)

	return router
}
