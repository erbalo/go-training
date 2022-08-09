package domain_test

import (
	"chapter3/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewOrderLine(t *testing.T) {
	orderLine := domain.NewOrderLine("tacos de pastor", 16, 15)

	assert.NotEmpty(t, orderLine)
	assert.Equal(t, int8(16), orderLine.Quantity)
	assert.Equal(t, float64(15), orderLine.UnitPrice)
}

func TestRetrieveAnIdentifier(t *testing.T) {
	orderLine := domain.NewOrderLine("tacos de suadero", 2, 10)

	assert.NotEmpty(t, orderLine)
	assert.NotEmpty(t, orderLine.GetId())
}
