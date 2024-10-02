package main

import (
	"errors"
	"fmt"
)

const USDEURKoeff = 0.84
const USDRUBKoeff = 92.0
const EURRUBKoeff = 1 / (USDEURKoeff / USDRUBKoeff)

func main() {
		for {
				converter()

				isRepeat := promptConvertAgain()

				if !isRepeat {
						break
				}
		}
}

func converter() {
	var initialCurrency string
	var targetCurrency string
	var amount float64
	var err error

	for {
		initialCurrency, err = requestInitialCurrency()
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	for {
		amount, err = requestAmount(initialCurrency)
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	for {
		targetCurrency, err = requestTargetCurrency(initialCurrency)
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	convert(initialCurrency, targetCurrency, amount)
}

func convert(initialCurrency string, targetCurrency string, amount float64) {
	var convertedAmount float64
	var err error

	switch initialCurrency {
	case "usd":
		convertedAmount, err = convertUSD(targetCurrency, amount)
	case "eur":
		convertedAmount, err = convertEUR(targetCurrency, amount)
	case "rub":
		convertedAmount, err = convertRUB(targetCurrency, amount)
	default:
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("%.2f %q = %.2f %q \n", amount, initialCurrency, convertedAmount, targetCurrency)
}

func convertRUB(targetCurrency string, amount float64) (float64, error) {
	switch targetCurrency {
	case "usd":
		return amount * (1 / USDRUBKoeff), nil
	case "eur":
		return amount * (1 / EURRUBKoeff), nil
	default:
		return 0, errors.New("Невозможно конвертировать")
	}
}

func convertEUR(targetCurrency string, amount float64) (float64, error) {
	switch targetCurrency {
	case "usd":
		return amount * (1 / USDEURKoeff), nil
	case "rub":
		return amount * EURRUBKoeff, nil
	default:
		return 0, errors.New("Невозможно конвертировать")
	}
}

func convertUSD(targetCurrency string, amount float64) (float64, error) {
	switch targetCurrency {
	case "eur":
		return amount * USDEURKoeff, nil
	case "rub":
		return amount * USDRUBKoeff, nil
	default:
		return 0, errors.New("Невозможно конвертировать")
	}
}

func requestAmount(currency string) (float64, error) {
	var amount float64
	fmt.Printf("Введите сумму для обмена в %q: ", currency)
	fmt.Scan(&amount)

	if amount <= 0 {
		return 0, errors.New("Неверно введна сумма")
	}

	return amount, nil
}

func requestInitialCurrency() (string, error) {
	var currency string
	fmt.Print("Введите исходную валюту (usd, eur, rub): ")
	fmt.Scan(&currency)

	if currency != "usd" && currency != "eur" && currency != "rub" {
		return "", errors.New("Неверный ввод")
	}

	return currency, nil
}

func requestTargetCurrency(initialCurrency string) (string, error) {
	var currency string
	fmt.Print("Введите целевую валюту (usd, eur, rub): ")
	fmt.Scan(&currency)

	if currency != "usd" && currency != "eur" && currency != "rub" {
		return "", errors.New("Неверный ввод")
	}

	if currency == initialCurrency {
		return "", errors.New("Целевая валюта не должна совпадать с начальной")
	}

	return currency, nil
}

func promptConvertAgain() bool {
		var answer string
		fmt.Print("Рассчитать другие данные (y/n)? ")
		fmt.Scan(&answer)

		return answer == "y" || answer == "Y"
}
