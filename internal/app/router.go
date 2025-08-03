package app

import (
	"github.com/go-chi/chi/v5"
	"go-subscription-service/internal/middleware"
)

func SetupRouter(container *AppContainer) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestLogger)

	container.SubscriptionHandler.RegisterRoutes(router)

	return router
}
