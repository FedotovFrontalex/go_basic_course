package file

import (
	"encoding/json"
	"errors"
	"jsonBin/print"
	"os"
)

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
	data, err := os.ReadFile(fileStorage.Filename)

	if err != nil {
		return nil, errors.New("Can't read file")
	}

	return data, nil
}

func (fileStorage *FileStorage) IsJson(name string) (bool, error) {
	data, err := fileStorage.Read()
	if err != nil {
		return false, err
	}
	isJson := json.Valid(data)

	return isJson, nil
}
