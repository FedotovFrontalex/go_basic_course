package storage

import (
	"encoding/json"
	"errors"
	"jsonBin/bins"
	"jsonBin/print"
	"time"
)

type Db interface {
	Write([]byte)
	Read() ([]byte, error)
}

type Storage struct {
	db Db
}

func NewStorage(db Db) *Storage {
	return &Storage{
		db: db,
	}
}

func (storage *Storage) SaveBinList(binList *bins.BinList) {
	binList.UpdatedAt = time.Now()
	data, err := binList.ToBytes()

	if err != nil {
		print.Error(err)
		return
	}
	storage.db.Write(data)
}

func (storage *Storage) AddBin(bin *bins.Bin) {
	list := storage.GetBinList()
	list.AddBin(*bin)
	storage.SaveBinList(list)
}

func (storage *Storage) DeleteBin(id string) {
	list := storage.GetBinList()
	list.DeleteBin(id)
	storage.SaveBinList(list)
}

func (storage *Storage) GetBinList() *bins.BinList {
	data, err := storage.db.Read()
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
