package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go-subscription-service/internal/model"
	"go-subscription-service/internal/service"
	"net/http"
)

type HTTPHandler struct {
	subscriptionService *service.SubscriptionService
}

func NewHTTPHandler(subSvc *service.SubscriptionService) *HTTPHandler {
	return &HTTPHandler{subscriptionService: subSvc}
}

func (h *HTTPHandler) RegisterRoutes(router *chi.Mux) {
	router.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", h.CreateSubscription)
		r.Get("/{id}", h.GetSubscription)
		r.Put("/{id}", h.UpdateSubscription)
		r.Delete("/{id}", h.DeleteSubscription)
		r.Get("/", h.ListSubscriptions)
	})
}

func (h *HTTPHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var sub model.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.subscriptionService.Create(r.Context(), &sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(sub)
	if err != nil {
		return
	}
}

func (h *HTTPHandler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	sub, err := h.subscriptionService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if sub == nil {
		http.Error(w, "Subscription not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sub); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *HTTPHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	var sub model.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sub.ID = id

	err = h.subscriptionService.Update(r.Context(), &sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(sub)
	if err != nil {
		return
	}
}

func (h *HTTPHandler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	err = h.subscriptionService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	subs, err := h.subscriptionService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(subs)
	if err != nil {
		return
	}
}
