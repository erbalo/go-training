package domain_test

import (
	"chapter3/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewOrderForUser(t *testing.T) {
	user := "Tony Stark"

	order := domain.NewOrder(user)

	assert.NotEmpty(t, order)
	assert.Equal(t, user, order.User)
	assert.Len(t, order.OrderLines, 0)
	assert.Equal(t, 0.0, order.TotalCost)
}

func TestShouldAppendNewOrderLineToOrder(t *testing.T) {
	user := "Dr. Strange"

	order := domain.NewOrder(user)
	orderLine := domain.NewOrderLine("tacos dorados", 3, 14.50)

	order.AppendLine(orderLine)

	assert.NotEmpty(t, order)
	assert.NotEmpty(t, orderLine)
	assert.Len(t, order.OrderLines, 1)
}

func TestShouldComputeCostGivingAnOrderLine(t *testing.T) {
	user := "Hulk"

	order := domain.NewOrder(user)
	orderLine1 := domain.NewOrderLine("tacos campechanos", 2, 10.50)
	orderLine2 := domain.NewOrderLine("coca-cola", 1, 20)

	order.AppendLine(orderLine1)
	order.AppendLine(orderLine2)
	order.ComputeCost()

	expectedCost := float64(orderLine1.Quantity) * orderLine1.UnitPrice
	expectedCost += float64(orderLine2.Quantity) * orderLine2.UnitPrice

	assert.NotEmpty(t, order)
	assert.NotEmpty(t, orderLine1)
	assert.NotEmpty(t, orderLine2)
	assert.Len(t, order.OrderLines, 2)
	assert.Equal(t, expectedCost, order.TotalCost)
}

func TestShouldRemoveAnOrderLine(t *testing.T) {
	user := "Thor"

	order := domain.NewOrder(user)
	orderLine1 := domain.NewOrderLine("tacos de longaniza", 5, 8.50)
	orderLine2 := domain.NewOrderLine("cerveza de barril", 1, 30.5)

	order.AppendLine(orderLine1)
	order.AppendLine(orderLine2)

	expectedCost := float64(orderLine2.Quantity) * orderLine2.UnitPrice

	order.RemoveLine(orderLine1)
	order.ComputeCost()

	assert.NotEmpty(t, order)
	assert.NotEmpty(t, orderLine1)
	assert.NotEmpty(t, orderLine2)
	assert.Len(t, order.OrderLines, 1)
	assert.Equal(t, expectedCost, order.TotalCost)
}

func TestShouldUpdatAnOrderLine(t *testing.T) {
	user := "Thor"

	order := domain.NewOrder(user)
	orderLine1 := domain.NewOrderLine("tacos de longaniza", 5, 8.50)
	orderLine2 := domain.NewOrderLine("cerveza de barril", 1, 30.5)

	order.AppendLine(orderLine1)
	order.AppendLine(orderLine2)

	orderLine1.Quantity = 8

	expectedCost := float64(orderLine1.Quantity) * orderLine1.UnitPrice
	expectedCost += float64(orderLine2.Quantity) * orderLine2.UnitPrice

	order.UpdateLine(orderLine1.GetId(), orderLine1)
	order.ComputeCost()

	assert.NotEmpty(t, order)
	assert.NotEmpty(t, orderLine1)
	assert.NotEmpty(t, orderLine2)
	assert.Len(t, order.OrderLines, 2)
	assert.Equal(t, int8(8), order.OrderLines[0].Quantity)
	assert.Equal(t, expectedCost, order.TotalCost)
}
