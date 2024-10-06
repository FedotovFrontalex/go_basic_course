package storage

import (
	"encoding/json"
	"errors"
	"jsonBin/bins"
	"jsonBin/file"
	"jsonBin/print"
	"time"
)

func SaveBinList(binList *bins.BinList) {
	binList.UpdatedAt = time.Now()
	data, err := binList.ToBytes()

	if err != nil {
		print.Error(err)
		return
	}
	file.WriteFile(data, "bins.json")
}

func GetBinList() *bins.BinList {
	data, err := file.ReadFile("bins.json")
	newBinList :=
		&bins.BinList{
			Bins:      []bins.Bin{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	if err != nil {
		print.Error(err)
		return newBinList
	}

	var binList bins.BinList

	err = json.Unmarshal(data, &binList)

	if err != nil {
		print.Error(errors.New("Can't unmarshal json"))
		return newBinList
	}

	return &binList
}
