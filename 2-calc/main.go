package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operations := []string{"AVG", "SUM", "MED"}
	operation := getOperation(operations)
	numbers := getNumbers()
	result := calculation(operation, numbers)

	fmt.Printf("Результат операции %s: %.2f\n", operation, result)
}

func getOperation(operations []string) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Введите желаемую операцию (%s): ", strings.Join(operations, ", "))
		scanner.Scan()
		input := strings.ToUpper(strings.TrimSpace(scanner.Text()))

		if slices.Contains(operations, input) {
			return input
		}
		fmt.Println("Ошибка: введите одну из доступных операций")
	}
}

func getNumbers() []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите числа через запятую: ")
		scanner.Scan()

		numbers, err := parseNumbers(scanner.Text())
		if err != nil {
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}

		if len(numbers) == 0 {
			fmt.Println("Ошибка: необходимо ввести хотя бы одно число")
			continue
		}

		return numbers
	}
}

func parseNumbers(str string) ([]float64, error) {
	var numbers []float64

	elements := strings.Split(str, ",")
	for index, value := range elements {
		value := strings.TrimSpace(value)
		if value == "" {
			continue
		}

		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("'%s' (элемент %d) не является числом", value, index+1)
		}
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		return nil, fmt.Errorf("не найдено ни одного числа")
	}
	return numbers, nil
}

func calculation(operation string, numbers []float64) float64 {
	var result float64
	switch operation {
	case "SUM":
		result = sum(numbers)
	case "AVG":
		result = avg(numbers)
	case "MED":
		result = med(numbers)
	default:
		return 0
	}
	return result
}

func sum(numbers []float64) float64 {
	var result float64
	for _, value := range numbers {
		result += value
	}
	return result
}

func avg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return sum(numbers) / float64(len(numbers))
}

func med(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)
	n := len(sorted)
	mid := n / 2
	if n%2 == 1 {
		return sorted[mid]
	}
	return (sorted[mid-1] + sorted[mid]) / 2
}
