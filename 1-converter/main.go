package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var validCurrencies = []string{"USD", "EUR", "RUB"}
	var exchangeRates = map[string]float64{
		"USD": 1.0,
		"EUR": 0.93,
		"RUB": 84.21,
	}

	originalCurrency := getAndValidateCurrency(validCurrencies, "исходную")
	remainingCurrencies := removeCurrency(validCurrencies, originalCurrency)

	if len(remainingCurrencies) == 0 {
		fmt.Println("Нет доступных валют для обмена")
		return
	}

	desiredCurrency := getAndValidateCurrency(remainingCurrencies, "желаемую")
	amount := getAndValidateAmount(originalCurrency)
	result := calculateExchange(originalCurrency, desiredCurrency, amount, exchangeRates)

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

func calculateExchange(from, to string, amount float64, exchangeRates map[string]float64) float64 {
	fromRate, fromExists := exchangeRates[from]
	toRate, toExists := exchangeRates[to]

	if !fromExists || !toExists {
		return 0
	}

	// Конвертируем через USD как промежуточную валюту
	amountInUSD := amount / fromRate
	return amountInUSD * toRate
}
