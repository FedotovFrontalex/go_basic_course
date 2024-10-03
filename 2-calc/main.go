package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var operation string
	var err error
	var result float64

	for {
		operation, err = requestOperationType()
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	operations := requestOpearations()
	result, err = calculate(operation, operations)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func requestOperationType() (string, error) {
	fmt.Println("Введите название операции")
	fmt.Println("sum - Посчитать сумму")
	fmt.Println("avg - Посчитать среднее")
	fmt.Println("med - Найти медиану")

	var userInput string
	fmt.Scan(&userInput)

	if userInput != "sum" && userInput != "avg" && userInput != "med" {
		return "", errors.New("Неверный тип операции")
	}

	return userInput, nil
}

func calculate(operationType string, operations []float64) (float64, error) {
	switch operationType {
	case "sum":
		return calcSum(operations), nil
	case "avg":
		return calcAVG(operations), nil
	case "med":
		return calcMed(operations), nil
	default:
		return 0.0, errors.New("Я не умею это считать")
	}
}

func requestOpearations() []float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите последовательность чисел через запятую. ")
	fmt.Println("Если введете 0 или значение отличное от числа, то такая операция не участвует в расчетах.")
	userInput, _ := reader.ReadString('\n')
	fmt.Println("Вы ввели: ", userInput)

	return normalizeData(userInput)
}

func normalizeData(initial string) []float64 {
	slice := strings.Split(initial, ",")
	numbers := []float64{}

	for index := range slice {
		num, err := trim(slice[index])
		if err == nil && num != 0 {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func trim(str string) (float64, error) {
	numStr := strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	num, err := strconv.ParseFloat(numStr, 64)

	return num, err
}

func calcSum(slice []float64) float64 {
	sum := 0.0
	for index := range slice {
		sum += slice[index]
	}

	return sum
}

func calcAVG(slice []float64) float64 {
	return calcSum(slice) / float64(len(slice))
}

func calcMed(slice []float64) float64 {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return slice[len(slice)/2]
}
