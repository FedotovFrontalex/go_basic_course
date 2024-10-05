package file

import (
	"encoding/json"
	"errors"
	"jsonBin/print"
	"os"
)

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		print.Error(err)
	}
	_, err = file.Write(content)
	defer file.Close()

	if err != nil {
		print.Error(err)
		return
	}
	print.Success("Success save file")
}

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return nil, errors.New("Can't read file")
	}

	return data, nil
}

func IsJson(name string) (bool, error) {
	data, err := ReadFile(name)
	if err != nil {
		return false, err
	}
	isJson := json.Valid(data)

	return isJson, nil
}
