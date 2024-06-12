package main

import (
	"log"
	"net/http"

	"github.com/Fred-devv/orders-service/http"
)

// Главная функция приложения
func main() {
	log.Println("Starting orders service...")

	// Подключение к БД
	db, err := utils.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Подключение к NATS Streaming
	natsConn, err := connetToNATS()
	if er != nil {
		log.Fatal(err)
	}
	defer natsConn.Close()

	// Создание репозитория заказов
	orderRepository := repositories.NewOrderRepository(db)

	// Создание сервиса заказов
	orderService := services.NewOrderService(orderRepository)

	// Создание обработчика заказов
	orderHandler := handlers.NewOrderHandler(orderService)

	// Создание HTTP-сервера
	httpServer := http.NewServer(orderHandler)

	// Запуск HTTP-сервера
	log.Fatal(httpSercer.ListenAndServe(":8080"))
}

//httpServer := http.NewServer()
//log.Fatal(httpServer.ListenAndServe(":8080"))
