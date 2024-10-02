package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("__Калькулятор индекса массы тела")
	for {
		userHeight, userWeight := getUserInputs()
		IMT, err := calculateIMT(userHeight, userWeight)

		if err != nil {
			//fmt.Println(err)
			//continue
			panic(err)
		}

		printResult(IMT)

		if !promtCalculate() {
			break
		}
	}
}

func calculateIMT(userHeight float64, userWeight float64) (float64, error) {
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("Не заданы рост или вес. Повторите ввод")
	}
	const IMTPower = 2
	return userWeight / math.Pow(userHeight/100, IMTPower), nil
}

func getUserInputs() (float64, float64) {
	var userHeight float64
	var userWeight float64

	fmt.Print("Введите ваш рост в см: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите ваш вес в кг: ")
	fmt.Scan(&userWeight)

	return userHeight, userWeight
}

func printResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f\n", IMT)

	switch {
	case IMT < 16:
		fmt.Println("У Вас сильный дефицит веса")
	case IMT < 18.5:
		fmt.Println("У Вас дефицит веса")
	case IMT < 25:
		fmt.Println("У Вас нормальный вес")
	case IMT < 30:
		fmt.Println("У Вас избыточный вес")
	default:
		fmt.Println("У Вас степень ожирения")
	}
}

func promtCalculate() bool {
	var answer string
	fmt.Print("Повторить расчет (y/n)?: ")
	fmt.Scanln(&answer)

	return answer == "y" || answer == "Y"
}
