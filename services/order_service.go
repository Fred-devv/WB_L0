package services

import (
	"github.com/Fred-devv/orders-service/models"
	"github.com/Fred-devv/orders-service/repositories"
)

// Сервис заказов
type OrderService struct {
	repository *repositories.OrderRepository
}

// Создание сервиса заказов
func NewOrderService(repository *repositories.OrderRepository) *OrderService {
	return &OrderService{repository: repository}
}

// Создание заказа с помощью репозитория
func (s *OrderService) CreateOrder(order *models.Order) error {
	// Реализация создания заказа с помощью репозитория
	return s.repository.CreateOrder(order)
}

// Получение заказа по id с помощью репозитория
func (s *OrderService) GetOrder(id int) (*models.Order, error) {
	// Релизация получения заказа по id с помощью репозитория
	return s.repository.GetOrder(id)
}
