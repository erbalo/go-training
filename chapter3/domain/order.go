package domain

import (
	"log"
	"math/rand"
	"time"
)

type Order struct {
	id         int64
	OrderLines []OrderLine
	User       string
	TotalCost  float64
	createdAt  time.Time
}

func NewOrder(user string) Order {
	id := rand.Int63()
	created := time.Now().UTC()

	return Order{
		id:         id,
		createdAt:  created,
		User:       user,
		OrderLines: []OrderLine{},
	}
}

func (order *Order) AppendLine(line OrderLine) {
	order.OrderLines = append(order.OrderLines, line)
}

func (order *Order) RemoveLine(line OrderLine) {
	lines := order.OrderLines

	for index, orderLine := range lines {
		if orderLine == line {
			order.OrderLines = append(lines[0:index], lines[index+1:]...)
			break
		}
	}
}

func (order *Order) UpdateLine(orderLineId int64, line OrderLine) {
	lines := order.OrderLines

	for index, orderLine := range lines {
		if orderLine.id == orderLineId {
			lines[index] = line
			order.OrderLines = lines
			break
		}
	}
}

func (order *Order) ComputeCost() {
	var total float64

	for _, orderLine := range order.OrderLines {
		total += orderLine.UnitPrice * float64(orderLine.Quantity)
	}

	order.TotalCost = total

	log.Printf("Cost: %2f\n", order.TotalCost)
}
