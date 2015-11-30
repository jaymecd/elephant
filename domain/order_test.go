package domain

import "testing"

func TestNewOrder(t *testing.T) {
	t.Log("New order")
	order := NewOrder(42)

	if id := order.Id(); id != 42 {
		t.Errorf("Expected Id of 42, but it was %d instead.", id)
	}
}
