package bins

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
	Bins []Bin
}

func (binList *BinList) AddBin(bin *Bin) {
	binList.Bins = append(binList.Bins, *bin)
}

func (binList *BinList) Print() {
	fmt.Println("Список bin")

	for index := range binList.Bins {
		binList.Bins[index].Print()
		fmt.Println("---------")
	}
}

func (bin *Bin) Print() {
	fmt.Println("id: ", bin.id)
	fmt.Println("private: ", bin.private)
	fmt.Println("createdAt: ", bin.createdAt)
	fmt.Println("name: ", bin.name)
}

func CreateBin(name string, binList *BinList) (*Bin, error) {
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

	binList.AddBin(bin)

	return bin, nil
}

func validateBinName(name string) error {
	if name == "" {
		return errors.New("Не передано название для Bin")
	}

	return nil
}
