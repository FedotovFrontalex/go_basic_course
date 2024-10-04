package main

import (
	"fmt"
	"jsonBin/bins"
)

func main() {
	var err error
	binList, err := initBinList()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := promptBinName()

	bin, err := bins.CreateBin(name, binList)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Новый bin")
	bin.Print()
	fmt.Println("----------------")
	binList.Print()
}

func initBinList() (*bins.BinList, error) {
	// Что-то делаем
	// Пока возвращаем пустой BinList

	return &bins.BinList{
		Bins: []bins.Bin{},
	}, nil
}

func promptBinName() string {
	var binName string
	fmt.Print("Введите название Bin: ")
	fmt.Scanln(&binName)
	return binName
}
