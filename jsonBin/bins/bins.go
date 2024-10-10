package bins

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins      []Bin     `json:"bins"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (binList *BinList) AddBin(bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}

func (binList *BinList) Print() {
	fmt.Println("Список bin")

	for index := range binList.Bins {
		binList.Bins[index].Print()
		fmt.Println("---------")
	}
}

func (bin *Bin) Print() {
	fmt.Println("id: ", bin.Id)
	fmt.Println("private: ", bin.Private)
	fmt.Println("createdAt: ", bin.CreatedAt)
	fmt.Println("name: ", bin.Name)
}

func CreateBin(name string, binList *BinList) error {
	err := validateBinName(name)
	if err != nil {
		return err
	}

	bin := Bin{
		Id:        "id будем получать из сервиса. Пока так",
		Private:   false,
		CreatedAt: time.Now(),
		Name:      name,
	}

	binList.AddBin(bin)

	return nil
}

func (binList *BinList) DeleteBin(id string) {
	var bins []Bin

	for _, value := range binList.Bins {
		if value.Id != id {
			bins = append(bins, value)
		}
	}

	binList.Bins = bins
}

func validateBinName(name string) error {
	if name == "" {
		return errors.New("bin name is required")
	}

	return nil
}

func (binList *BinList) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(binList)

	if err != nil {
		return nil, errors.New("Can't create json")
	}

	return bytes, nil
}
