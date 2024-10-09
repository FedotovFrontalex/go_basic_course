package testHelpers

import (
	"errors"
	"os"
)

func CreateMockFile(fileType string) (string, error) {
	var filename string
	var content []byte
	switch fileType {
	case "txt":
		filename, content = createMockTxt()
	case "json":
		filename, content = createMockJson()
	default:
		return "", errors.New("Can't create file")
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return "", err
	}
	_, err = file.Write(content)
	if err != nil {
		return "", nil
	}
	return filename, nil
}

func createMockJson() (string, []byte) {
	filename := "mock.json"
	content := []byte("{\"test\": 5}")
	return filename, content
}

func createMockTxt() (string, []byte) {
	filename := "mock.txt"
	content := []byte("validContent")
	return filename, content
}
