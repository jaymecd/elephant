package main

import (
	"fmt"
	"github.com/jaymecd/elephant/domain/order"
)

func main() {
	order := NewOrder(123)

	fmt.Printf("Order ID : %d\n", order.Id())
}
