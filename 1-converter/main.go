package main

import "fmt"

func main() {
	const USDEURKoeff = 0.84
	const USDRUBKoeff = 92.0
	const EURRUBKoeff = 1 / (USDEURKoeff / USDRUBKoeff)

	fmt.Println("Конвертация EUR в RUB по коэффициенту")
	fmt.Println(EURRUBKoeff)

	fmt.Println("---------------------------------")

	fmt.Print("1р = ")
	fmt.Print(1 / USDRUBKoeff)
	fmt.Println("usd")
	
	fmt.Print("1eur = ")
	fmt.Print(1 / USDEURKoeff)
	fmt.Println("usd")
	
	fmt.Print("1eur = ")
	fmt.Print(1 * EURRUBKoeff)
	fmt.Println("р")
}
