package main

import (
	"fmt"
	"github.com/jaymecd/elephant/domain"
)

var (
	version   = "undefined"
	buildTime = "undefined"
	gitHash   = "undefined"
)

func main() {
	fmt.Printf("Version    : %s\n", version)
	fmt.Printf("Git Hash   : %s\n", gitHash)
	fmt.Printf("Build Time : %s\n", buildTime)

	order := domain.NewOrder(123)

	fmt.Printf("\nOrder ID : %d\n", order.Id())
}
