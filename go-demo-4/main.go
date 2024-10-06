package main

import (
	"demo/app-4/account"
	jsonDb "demo/app-4/jsonDB"
	"demo/app-4/print"
	"errors"
	"fmt"
)

func main() {
	db := jsonDb.NewJsonDb("accounts.json")
	vault := account.NewVault(db)

	for {
		userInput := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите пункт меню",
		})

		if userInput == "4" {
			break
		}

		switch userInput {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			print.Error(errors.New("Неверный пункт меню"))
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	link := promptData([]string{"Введите url"})

	account1, err := account.NewAccount(login, password, link)

	if err != nil {
		print.Error(err)
		return
	}

	vault.AddAccount(*account1)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url"})

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

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url"})
	err := vault.DeleteAccountByUrl(url)

	if err != nil {
		print.Error(err)
		return
	}

	print.Success("Операция выполнена")
}

func promptData[T any](content []T) string {
	var res string
	for idx := range content {
		if idx != len(content)-1 {
			print.Prompt(content[idx], true)
			continue
		}
		print.Prompt(content[idx], false)
		print.Prompt(": ", false)
	}
	fmt.Scanln(&res)
	return res
}
