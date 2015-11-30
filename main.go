package main

import (
	"fmt"
	"github.com/jaymecd/elephant/domain"
)

func main() {
	order := domain.NewOrder(123)

	fmt.Printf("Order ID : %d\n", order.Id())
}
