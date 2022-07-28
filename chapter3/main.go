package main

import (
	"chapter3/domain"
	"encoding/json"
	"log"
)

func main() {
	order := domain.NewOrder("erick")
	orderLine1 := domain.NewOrderLine("churritos", 1, 1)
	orderLine2 := domain.NewOrderLine("churritos picosos", 1, 2)
	orderLine3 := domain.NewOrderLine("churritos de amaranto", 1, 3)
	orderLine4 := domain.NewOrderLine("churritos mega picosos", 1, 4)
	order.AppendLine(orderLine1)
	order.AppendLine(orderLine2)
	order.AppendLine(orderLine3)
	order.AppendLine(orderLine4)

	orderLine1.Item = "pepe charros"
	order.UpdateLine(orderLine1.GetId(), orderLine1)
	order.RemoveLine(orderLine2)

	data, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", data)

	order.ComputeCost()
}
