package models

import (
	"errors"
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *Order) Valodate() error {
	if o.OrderID == "" {
		return errors.New("order_id is required")
	}
	if len(o.Data) == 0 {
		return errors.New("data is required")
	}
}
