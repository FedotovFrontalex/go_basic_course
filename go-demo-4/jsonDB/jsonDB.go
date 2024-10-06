package jsonDb

import (
	"demo/app-4/print"
	"os"
)

type JsonDb struct {
		Filename string
}

func NewJsonDb(filename string) *JsonDb {
		return &JsonDb{
				Filename: filename,
		}
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.Filename)

	if err != nil {
		print.Error(err)
	}
	_, err = file.Write(content)
	defer file.Close()

	if err != nil {
		print.Error(err)
		return
	}
	print.Success("Запись успешна")
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.Filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}
