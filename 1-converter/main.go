package main

import (
	"errors"
	"fmt"
)

const USDEURKoeff = 0.84
const USDRUBKoeff = 92.0
const EURRUBKoeff = 1 / (USDEURKoeff / USDRUBKoeff)

type currencyMap = map[string]float64
type baseCurrencyMap = map[string]currencyMap

func main() {
	currencyKoeffs := initKoeff()

	for {
		converter(currencyKoeffs)

		isRepeat := promptConvertAgain()

		if !isRepeat {
			break
		}
	}
}

func initKoeff() baseCurrencyMap {
	eurMap := currencyMap{
		"usd": 1 / USDEURKoeff,
		"rub": EURRUBKoeff,
	}
	usdMap := currencyMap{
		"eur": USDEURKoeff,
		"rub": USDRUBKoeff,
	}

	rubMap := currencyMap{
		"eur": 1 / EURRUBKoeff,
		"usd": 1 / USDRUBKoeff,
	}

	baseCurrency := baseCurrencyMap{
		"usd": usdMap,
		"eur": eurMap,
		"rub": rubMap,
	}

	return baseCurrency
}

func converter(koeffs baseCurrencyMap) {
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

	convert(koeffs, initialCurrency, targetCurrency, amount)
}

func convert(koeffs baseCurrencyMap, initialCurrency string, targetCurrency string, amount float64) {
	convertedAmount := amount * (koeffs)[initialCurrency][targetCurrency]
	fmt.Printf("%.2f %q = %.2f %q \n", amount, initialCurrency, convertedAmount, targetCurrency)
}

func requestAmount(currency string) (float64, error) {
	var amount float64
	fmt.Printf("Введите сумму для обмена в %q: ", currency)
	fmt.Scan(&amount)

	if amount <= 0 {
		return 0, errors.New("Неверная сумма")
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
