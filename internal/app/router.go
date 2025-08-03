package app

import "github.com/go-chi/chi/v5"

func SetupRouter(container *AppContainer) *chi.Mux {
	router := chi.NewRouter()

	container.SubscriptionHandler.RegisterRoutes(router)

	return router
}
