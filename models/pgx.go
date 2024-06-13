package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

func Connect() (*sql.DB, error) {
	// Подключение к PostgreSQL
	conn, err := sql.Open("postgres", "user:password@localhost:5432/orders_db")
	if err != nil {
		log.Fatal(err)
	}
	return conn, nil
}

func SaveOrder(conn *sql.DB, order Order) error {
	// Сохранение заказа в БД
	_, err := conn.Exec(context.Background(), "INSERT INTO orders (order_id, data) VALUES ($1, $2)", order.OrderID, order.Data)
	if err != nil {
		log.Println(err)
	}
	return err
}

func RestoreCache(conn *sql.DB) {
	// Восстановление кэша из БД
	rows, err := conn.Query(context.Background(), "SELECT order_id, data FROM orders")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		var orderID string
		var data []byte
		if err := rows.Scan(&orderID, &data); err != nil {
			log.Println(err)
			continue
		}
		if err := json.Unmarshal(data, &order); err != nil {
			log.Println(err)
			continue
		}
		cache[orderID] = order
	}
}
