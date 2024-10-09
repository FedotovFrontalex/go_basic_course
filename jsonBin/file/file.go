package file

import (
	"encoding/json"
	"errors"
	"jsonBin/print"
	"os"
)

var ErrReadFile = errors.New("Can't read file")
var ErrNotJson = errors.New("is not json")

type FileStorage struct {
	Filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{
		Filename: filename,
	}
}

func (fileStorage *FileStorage) Write(content []byte) {
	file, err := os.Create(fileStorage.Filename)

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

func (fileStorage *FileStorage) Read() ([]byte, error) {
	print.Message(fileStorage.Filename)
	data, err := os.ReadFile(fileStorage.Filename)

	if err != nil {
		return nil, errors.New("Can't read file")
	}

	return data, nil
}

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, ErrReadFile 	
}

	isJsonData := isJson(data)

	if !isJsonData {
		return nil, ErrNotJson 
	}
	return data, nil
}

func isJson(data []byte) bool {
	isJson := json.Valid(data)

	return isJson
}
