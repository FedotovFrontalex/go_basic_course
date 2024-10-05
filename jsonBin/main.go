package main

import (
	"fmt"
	"jsonBin/bins"
	"jsonBin/print"
	"jsonBin/storage"
)

func main() {
	var err error
	binList := storage.GetBinList()

	name := promptBinName()

	err = bins.CreateBin(name, binList)

	if err != nil {
		fmt.Println(err)
		return
	}

	storage.SaveBinList(binList)
}

func promptBinName() string {
	var binName string
	print.Prompt("Enter bin name: ", false)
	fmt.Scanln(&binName)
	return binName
}
