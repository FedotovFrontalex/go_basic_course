package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func (binList *BinList) addBin(bin *Bin) {
	binList.bins = append(binList.bins, *bin)
}

func (binList *BinList) print() {
	fmt.Println("Список bin")

	for index := range binList.bins {
		binList.bins[index].print()
		fmt.Println("---------")
	}
}

func (bin *Bin) print() {
	fmt.Println("id: ", bin.id)
	fmt.Println("private: ", bin.private)
	fmt.Println("createdAt: ", bin.createdAt)
	fmt.Println("name: ", bin.name)
}

func main() {
	var err error
	binList, err := initBinList()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := promptBinName()

	bin, err := createBin(name, binList)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Новый bin")
	bin.print()
	fmt.Println("----------------")
	binList.print()
}

func initBinList() (*BinList, error) {
	// Что-то делаем
	// Пока возвращаем пустой BinList

	return &BinList{
		bins: []Bin{},
	}, nil
}

func promptBinName() string {
	var binName string
	fmt.Print("Введите название Bin: ")
	fmt.Scanln(&binName)
	return binName
}

func createBin(name string, binList *BinList) (*Bin, error) {
	err := validateBinName(name)
	if err != nil {
		return nil, err
	}

	bin := &Bin{
		id:        "id будем получать из сервиса. Пока так",
		private:   false,
		createdAt: time.Now(),
		name:      name,
	}

	binList.addBin(bin)

	return bin, nil
}

func validateBinName(name string) error {
	if name == "" {
		return errors.New("Не передано название для Bin")
	}

	return nil
}
