package domain

type Order struct {
	id int
}

// Provides new instance of Order
func NewOrder(id int) (*Order) {
	order := new(Order)
	order.id = id
	return order
}

// Provides ID of an order
func (this *Order) Id() (int) {
	return this.id
}
