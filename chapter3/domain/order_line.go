package domain

import (
	"math/rand"
	"time"
)

type OrderLine struct {
	id        int64
	createdAt time.Time
	Item      string
	Quantity  int8
	UnitPrice float64
}

// TODO assuming no validation errors
func NewOrderLine(item string, quantity int8, price float64) OrderLine {
	id := rand.Int63()
	created := time.Now().UTC()

	return OrderLine{
		id:        id,
		createdAt: created,
		Item:      item,
		Quantity:  quantity,
		UnitPrice: price,
	}
}

func (order *OrderLine) GetId() int64 {
	return order.id
}
