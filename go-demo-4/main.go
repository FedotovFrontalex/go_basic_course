package main

import (
	"demo/app-4/account"
	"demo/app-4/jsonDB"
	"demo/app-4/print"
	"errors"
	"fmt"
	"strings"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount(checkUrl),
	"3": findAccount(checkLogin),
	"4": deleteAccount,
}

func main() {
	db := jsonDb.NewJsonDb("accounts.json")
	vault := account.NewVault(db)

	for {
		userInput := promptData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по url",
			"3. Найти аккаунт по login",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите пункт меню",
		)

		if userInput == "5" {
			break
		}

		if menu[userInput] == nil {
			print.Error(errors.New("Неверный пункт меню"))
			continue
		}

		menu[userInput](vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	link := promptData("Введите url")

	account1, err := account.NewAccount(login, password, link)

	if err != nil {
		print.Error(err)
		return
	}

	vault.AddAccount(*account1)
}
func checkUrl(account account.Account, url string) bool {
	return strings.Contains(account.Link, url)
}

func checkLogin(account account.Account, login string) bool {
	return strings.Contains(account.Login, login)
}

func findAccount(checkFn func(account.Account, string) bool) func(*account.VaultWithDb) {
	return func(vault *account.VaultWithDb) {
		url := promptData("Введите для поиска")

		accounts := vault.FindAccounts(url, checkFn)

		if len(accounts) == 0 {
			print.Error(errors.New("Аккаунт не найден"))
		}

		for idx, account := range accounts {
			print.Message(idx + 1)
			account.Print()
			print.Message("-----------")
		}
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	err := vault.DeleteAccountByUrl(url)

	if err != nil {
		print.Error(err)
		return
	}

	print.Success("Операция выполнена")
}

func promptData(content ...any) string {
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
