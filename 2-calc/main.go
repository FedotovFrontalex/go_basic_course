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

var operationsMap = map[string]func([]float64) float64{
	"sum": calcSum,
	"avg": calcAVG,
	"med": calcMed,
}

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

	operations := requestOperations()
	result, err = calculate(operation, operations)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func requestOperationType() (string, error) {
	userInput := promptData(
		"sum - Calculate the amount",
		"avg - Calculate the average",
		"med - Calculate the median",
		"Enter the name of the operation",
	)

	if userInput != "sum" && userInput != "avg" && userInput != "med" {
		return "", errors.New("Invalid operation type")
	}

	return userInput, nil
}

func calculate(operationType string, operations []float64) (float64, error) {
	operationFn := operationsMap[operationType]

	if operationFn == nil {
		return 0.0, errors.New("I can't do it.")
	}

	return operationFn(operations), nil
}

func requestOperations() []float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sequence of numbers separated by commas.")
	fmt.Println("If you enter 0 or a value other than a number, then this operation is not involved in the calculations.")
	userInput, _ := reader.ReadString('\n')

	return normalizeData(userInput)
}

func normalizeData(initial string) []float64 {
	slice := strings.Split(initial, ",")
	numbers := []float64{}

	for index := range slice {
		num, err := trim(slice[index])
		if err == nil && num != 0.0 {
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

	if len(slice)%2 != 0 {
		return slice[len(slice)/2]
	} else {
		idx := len(slice) / 2
		return (slice[idx] + slice[idx-1]) / 2
	}

}

func promptData(content ...any) string {
	var res string
	for idx := range content {
		if idx != len(content)-1 {
			fmt.Println(content[idx])
			continue
		}
		fmt.Printf("%v: ", content[idx])
	}
	fmt.Scanln(&res)
	return res
}
