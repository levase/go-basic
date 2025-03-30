package main

import (
	"fmt"
	"slices"
	"strings"
)

const (
	USDToEUR = 0.93
	USDToRUB = 84.21
)

func main() {
	var validCurrencies = []string{"USD", "EUR", "RUB"}

	originalCurrency := getAndValidateCurrency(validCurrencies, "исходную")
	remainingCurrencies := removeCurrency(validCurrencies, originalCurrency)
	
	if len(remainingCurrencies) == 0 {
		fmt.Println("Нет доступных валют для обмена")
		return
	}

	desiredCurrency := getAndValidateCurrency(remainingCurrencies, "желаемую")
	amount := getAndValidateAmount(originalCurrency)
	result := calculateExchange(originalCurrency, desiredCurrency, amount)

	fmt.Printf("\nРезультат обмена: %.2f %s = %.2f %s\n", amount, originalCurrency, result, desiredCurrency)
}

func clearScanBuffer() {
	var discard string
	fmt.Scanln(&discard)
}

func getCurrency(validCurrencies []string, step string) string {
	var currency string

	fmt.Printf("Введите "+step+" валюту (%s): ", strings.Join(validCurrencies, ", "))
	_, err := fmt.Scan(&currency)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		clearScanBuffer()
	}
	return currency
}

func getAndValidateCurrency(validCurrencies []string, step string) string {
	for {
		currency := getCurrency(validCurrencies, step)
		currency = strings.ToUpper(currency)

		if slices.Contains(validCurrencies, currency) {
			return currency
		}
		fmt.Println("Пожалуйста, введите корректное название валюты")
	}
}

func removeCurrency(slice []string, currency string) []string {
	result := make([]string, 0)
	for _, v := range slice {
		if v == currency {
			continue
		}
		result = append(result, v)
	}
	return result
}

func getAmount(currency string) float64 {
	var amount float64

	fmt.Printf("Введите количество %s для обмена: ", currency)
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		clearScanBuffer()
	}
	return amount
}

func getAndValidateAmount(currency string) float64 {
	for {
		amount := getAmount(currency)

		if amount > 0 {
			return amount
		}
		fmt.Println("Пожалуйста, введите корректное количество валюты")
	}
}

func calculateExchange(from, to string, amount float64) float64 {
	var inUSD float64

	// Конвертируем все в USD как промежуточную валюту
	switch from {
	case "USD":
		inUSD = amount
	case "EUR":
		inUSD = amount / USDToEUR
	case "RUB":
		inUSD = amount / USDToRUB
	}

	// Конвертируем из USD в целевую валюту
	switch to {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * USDToEUR
	case "RUB":
		return inUSD * USDToRUB
	default:
		return 0
	}
}
