package models

import (
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
)

func Subscribe(sc *stan.Conn) {
	// Подписка на канал NATS Streaming
	sub, err := sc.Subscribe("orders", func(msg *stan.Msg) {
		// Обработка полученного сообщения
		var order Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Println(err)
			return
		}

		// Сохранение заказа в кэш
		cache[order.ID] = order
	}, stan.DurableName("order-service"))
	if err != nil {
		log.Println(err)
	}
	defer sub.Unsubscribe()
}

var cache = make(map[string]Order)
