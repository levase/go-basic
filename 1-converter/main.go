package main

import "fmt"

func main() {
	const USDToEUR = 0.93
	const USDToRUB = 84.21

	// Рассчитываем, сколько рублей в 1 евро
	EURToRUB := (1 / USDToEUR) * USDToRUB

	fmt.Printf("1 EUR = %.2f RUB\n", EURToRUB)
}
