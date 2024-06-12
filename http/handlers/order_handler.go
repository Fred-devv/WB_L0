package handlers

import (
	"encoding/json"
	"net/http"
	"projects/wb_L0/services"
	"strconv"
	"strings"
)

// Обработчик заказов
type OrderHandler struct {
	service *services.OrderService
}

// Создание обработчика заказов
func NewOrderHandler(services *services.OrderService) *OrderHandler {
	return &OrderHandler{service: services}
}

// Обработка GET-запроса для получения заказа по id
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	// Реализация обработчика GET - запроса для получения заказа по id
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/orders"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := h.service.GetOrder(id)
	if err != nill {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(order)
}
