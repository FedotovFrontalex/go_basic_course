package main

import (
	"errors"
	"flag"
	"jsonBin/api"
	"jsonBin/config"
	"jsonBin/file"
	"jsonBin/print"
	"jsonBin/storage"
)

func main() {
	binsFileStorage := file.NewFileStorage("bins.json")
	localStorage := storage.NewStorage(binsFileStorage)

	filename := flag.String("file", "", "filename")
	name := flag.String("name", "", "bin name")
	isCreate := flag.Bool("create", false, "create bin mode")
	id := flag.String("id", "", "bin id")
	isUpdate := flag.Bool("update", false, "update bin mode")
	isGet := flag.Bool("get", false, "get bin mode")
	isDelete := flag.Bool("delete", false, "delete bin mode")
	isList := flag.Bool("list", false, "list bin mode")

	flag.Parse()

	if *isCreate {
		createBin(*filename, *name, *localStorage)
		return
	}

	if *isUpdate {
		updateBin(*filename, *id)
		return
	}

	if *isGet {
		getBin(*id)
		return
	}

	if *isDelete {
		deleteBin(*id)
		return
	}

	if *isList {
		printBinsFromLocal(*localStorage)
		return
	}

	print.Error(errors.New("Invalid Parameters"))
}

func createBin(filename string, binName string, localStorage storage.Storage) {
	if filename == "" || binName == "" {
		print.Error(errors.New("No filename or bin name provided"))
		return
	}

	configApi := config.NewConfig()
	binApi := api.Init(configApi)
	bin, err := binApi.Create(filename, binName)

	if err != nil {
		print.Error(err)
		return
	}

	localStorage.AddBin(bin)
	print.Success("create successfully")
}

func updateBin(filename string, id string) {
	if filename == "" || id == "" {
		print.Error(errors.New("No filename or bin id provided"))
		return
	}
	configApi := config.NewConfig()
	binApi := api.Init(configApi)
	err := binApi.Update(filename, id)

	if err != nil {
		print.Error(err)
		return
	}

	print.Success("update successfully")
}

func getBin(id string) {
	if id == "" {
		print.Error(errors.New("no id provided"))
		return
	}

	configApi := config.NewConfig()
	binApi := api.Init(configApi)
	data, err := binApi.Get(id)

	if err != nil {
		print.Error(err)
		return
	}

	print.Message(data)
}

func deleteBin(id string) {
	if id == "" {
		print.Error(errors.New("no id provided"))
	}

	configApi := config.NewConfig()
	binApi := api.Init(configApi)
	err := binApi.Delete(id)

	if err != nil {
		print.Error(err)
		return
	}

	print.Success("delete successfully")
}

func printBinsFromLocal(localStorage storage.Storage) {
	binList := localStorage.GetBinList()
	binList.Print()
}
