package main

import (
	"errors"
	"fmt"
)

type bookmarkMap = map[string]string

func main() {
	var userInput int8
	var err error
	bookmarks := bookmarkMap{}
	fmt.Println("__Хранилище закладок__")

	for {
		userInput, err = printMenu()

		if err != nil {
			fmt.Println(err)
			continue
		}

		if userInput == 4 {
			fmt.Println("Хорошего дня")
			break
		}

		_, err := doAction(userInput, bookmarks)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func printMenu() (int8, error) {
	var userInput int8

	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")

	fmt.Scan(&userInput)

	if userInput <= 0 || userInput > 4 {
		return 0, errors.New("Неверный пункт меню")
	}

	return userInput, nil
}

func doAction(actionType int8, m bookmarkMap) (bool, error) {
	switch actionType {
	case 1:
		printBookmarks(m)
	case 2:
		addBookmark(m)
	case 3:
		deleteBookmark(m)
	default:
		return false, errors.New("Неизвестная операция")
	}

	return true, nil
}

func printBookmarks(m bookmarkMap) {
	for key, value := range m {
		fmt.Printf("%s - %s\n", key, value)
	}
}

func addBookmark(m bookmarkMap) {
	name := requestNameBookmark()
	url := requestUrlBookmark()
	m[name] = url
	fmt.Println("Закладка добавлена")
}

func deleteBookmark(m bookmarkMap) {
	name := requestNameBookmark()
	delete(m, name)
	fmt.Println("Закладка удалена")
}

func requestNameBookmark() string {
	var name string
	fmt.Print("Название закладки: ")
	fmt.Scan(&name)
	return name
}

func requestUrlBookmark() string {
	var url string
	fmt.Print("url закладки: ")
	fmt.Scan(&url)
	// Можно добавить валидацию
	return url
}
