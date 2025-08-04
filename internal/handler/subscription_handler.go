package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go-subscription-service/internal/model"
	"go-subscription-service/internal/service"
	"net/http"
	"time"
)

type HTTPHandler struct {
	subscriptionService *service.SubscriptionService
}

func NewHTTPHandler(subSvc *service.SubscriptionService) *HTTPHandler {
	return &HTTPHandler{subscriptionService: subSvc}
}

func (h *HTTPHandler) RegisterRoutes(router *chi.Mux) {
	router.Route("/subscriptions", func(r chi.Router) {
		// CRUD
		r.Post("/", h.CreateSubscription)
		r.Get("/{id}", h.GetSubscription)
		r.Put("/{id}", h.UpdateSubscription)
		r.Delete("/{id}", h.DeleteSubscription)
		r.Get("/", h.ListSubscriptions)

		// CUSTOM
		r.Get("/subscriptions/total", h.GetSubscriptionsTotal)
	})
}

// CreateSubscription godoc
// @Summary Создать новую подписку
// @Description Создает запись о подписке пользователя
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body model.Subscription true "Данные подписки"
// @Success 201 {object} model.Subscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
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

// GetSubscription godoc
// @Summary Получить подписку по ID
// @Description Возвращает информацию о подписке по её ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "UUID подписки"
// @Success 200 {object} model.Subscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [get]
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

// UpdateSubscription godoc
// @Summary Обновить подписку
// @Description Изменяет данные подписки по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "UUID подписки"
// @Param subscription body model.Subscription true "Данные подписки"
// @Success 200 {object} model.Subscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [put]
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

// DeleteSubscription godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "UUID подписки"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [delete]
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

// ListSubscriptions godoc
// @Summary Получить список подписок
// @Description Возвращает список всех подписок
// @Tags subscriptions
// @Accept json
// @Produce json
// @Success 200 {array} model.Subscription
// @Failure 500 {object} map[string]string
// @Router /subscriptions [get]
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

// GetSubscriptionsTotal godoc
// @Summary Получить суммарную стоимость подписок за период
// @Description Считает сумму подписок за выбранный период с фильтрами по user_id и service_name
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id query string false "UUID пользователя"
// @Param service_name query string false "Название сервиса"
// @Param from query string true "Дата начала периода (YYYY-MM-DD)"
// @Param to query string true "Дата конца периода (YYYY-MM-DD)"
// @Success 200 {object} map[string]int "total"
// @Failure 400 {object} map[string]string
// @Router /subscriptions/total [get]
func (h *HTTPHandler) GetSubscriptionsTotal(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		http.Error(w, "Invalid 'from' date", http.StatusBadRequest)
		return
	}
	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		http.Error(w, "Invalid 'to' date", http.StatusBadRequest)
		return
	}

	var userID *uuid.UUID
	if userIDStr != "" {
		uid, err := uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, "Invalid 'user_id'", http.StatusBadRequest)
			return
		}
		userID = &uid
	}

	total, err := h.subscriptionService.TotalPrice(r.Context(), userID, serviceName, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]int{"total": total})
}
