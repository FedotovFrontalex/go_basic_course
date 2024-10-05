package main

import (
	"demo/app-4/account"
	"demo/app-4/print"
	"errors"
	"fmt"
)

func main() {
	vault := account.NewVault()

	for {
		menu, err := getMenu()
		if err != nil {
			print.Error(err)
			continue
		}

		if menu == 4 {
			break
		}

		switch menu {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			print.Error(errors.New("Неверный пункт меню"))
		}
	}
}

func getMenu() (int8, error) {
	var userInput int8
	print.Message("Выберите пункт меню")
	print.Prompt("1. Создать аккаунт", true)
	print.Prompt("2. Найти аккаунт", true)
	print.Prompt("3. Удалить аккаунт", true)
	print.Prompt("4. Выход", true)
	fmt.Scanln(&userInput)

	if userInput < 1 && userInput > 4 {
		return 0, errors.New("Неверный пункт меню")
	}

	return userInput, nil
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	link := promptData("Введите url: ")

	account1, err := account.NewAccount(login, password, link)

	if err != nil {
		print.Error(err)
		return
	}

	vault.AddAccount(*account1)
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите url: ")

	accounts := vault.FindAccountsByUrl(url)

	if len(accounts) == 0 {
		print.Error(errors.New("Аккаунт не найден"))
	}

	for idx, account := range accounts {
		print.Message(idx + 1)
		account.Print()
		print.Message("-----------")
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите url: ")
	err := vault.DeleteAccountByUrl(url)

	if err != nil {
		print.Error(err)
		return
	}

	print.Success("Операция выполнена")
}

func promptData(text string) string {
	var res string
	print.Prompt(text, false)
	fmt.Scanln(&res)
	return res
}
