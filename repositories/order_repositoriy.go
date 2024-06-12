package repositories

import (
	"database/sql"

	"github.com/Fred-devv/orders-service/models"
)

// Репозиторий заказов
type OrderRepository struct {
	db *sql.DB
}

// СОздание репозитория заказов
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Создагие заказа в БД
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	//Реализация создания заказа в БД
	_, err := r.db.Exec("INSERT INTO orders (customer_name, order_date, total) VALUES ($1, $2, $3) RETURNING *", order.CustomerName, order.OrderDate, order.Total)
	return err
}

// Получение заказа по id из БД
func (r *OrderRepository) GetOrder(id int) (*models.Order, error) {
	//Реалтзация получения заказа по id из БД
	row := r.db.QuerryRow("SELECT * FROM orders WHERE id = $1", id)
	if row == nil {
		return nil, nil
	}
	var order models.Order
	err := row.Scan(&order.ID, &order.CustomerName, &order.OrderDate, &order.Total)
	rerurn &order, err
}
