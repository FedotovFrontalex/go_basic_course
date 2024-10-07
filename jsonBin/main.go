package main

import (
	"fmt"
	"jsonBin/api"
	"jsonBin/bins"
	"jsonBin/config"
	"jsonBin/file"
	"jsonBin/print"
	"jsonBin/storage"
)

func main() {
	var err error
	config := config.NewConfig()
	api.Init(config)
	fileStorage := file.NewFileStorage("bins.json")
	binStorage := storage.NewStorage(fileStorage)

	binList := binStorage.GetBinList()

	name := promptBinName()

	err = bins.CreateBin(name, binList)

	if err != nil {
		fmt.Println(err)
		return
	}

	binStorage.SaveBinList(binList)
}

func promptBinName() string {
	var binName string
	print.Prompt("Enter bin name: ", false)
	fmt.Scanln(&binName)
	return binName
}
