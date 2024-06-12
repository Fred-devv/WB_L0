package models

// Модель заказа
type Order struct {
	ID           int
	CustomerName string
	OrderDate    string
	Total        float64
}
