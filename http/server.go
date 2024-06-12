package http

import (
	"net/http"

	"github.com/Fred-devv/orders-service/http/handlers"
)

func NewServer(handler *handlers.OrderHandler) *http.Server {
	// Реализация создания http-сервера
}
