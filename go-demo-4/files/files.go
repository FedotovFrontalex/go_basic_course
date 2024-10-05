package files

import (
	"demo/app-4/print"
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
	print.Success("Запись успешна")
}

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return data, nil
}
