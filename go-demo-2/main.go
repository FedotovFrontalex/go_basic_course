package main

import (
	"fmt"
)

func main() {
	fmt.Println("__Калькулятор баланса__")

	transactions := getTransactions()
	sum := summTransactions(transactions)

	fmt.Printf("Ваш баланс: %.2f\n", sum)
}

func getTransactions() []float64 {
	transactions := []float64{}

	for {
		transaction := requestTransaction()

		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}

func requestTransaction() float64 {
	var transaction float64
	fmt.Print("Введите сумму транзакции. Для окончания ввода введите n: ")
	fmt.Scan(&transaction)

	return transaction
}

func summTransactions(transactions []float64) float64 {
	sum := 0.0

	for index := range transactions {
		sum += transactions[index]
	}

	return sum
}
