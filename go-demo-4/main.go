package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var passwordChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

type account struct {
	login    string
	password string
	link     string
}

type accountWithTimeStamp struct {
	account
	createAt  time.Time
	updatedAt time.Time
}

func (acc *account) print() {
	fmt.Printf("Логин - %s\n", acc.login)
	fmt.Printf("Пароль - %s\n", acc.password)
	fmt.Printf("Ссылка - %s\n", acc.link)
}

func (acc *accountWithTimeStamp) print() {
	fmt.Printf("Создан: %s\n", acc.createAt)
	fmt.Printf("Обновлен: %s\n", acc.updatedAt)
	acc.account.print()
}

func (acc *account) generatePassword(n int) {
	pass := make([]rune, n)
	lenghtChars := len(passwordChars)

	for index := range pass {
		pass[index] = passwordChars[rand.IntN(lenghtChars)]
	}

	acc.password = string(pass)
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	link := promptData("Введите url: ")

	account1, err := newAccountWithTimeStamp(login, password, link)

	if err != nil {
		fmt.Println(err)
		return
	}

	account1.print()
}

func promptData(text string) string {
	var res string
	fmt.Print(text)
	fmt.Scanln(&res)
	return res
}

func newAccountWithTimeStamp(login string, password string, link string) (*accountWithTimeStamp, error) {
	var err error

	err = validateLogin(login)
	if err != nil {
		return nil, err
	}

	err = validateUrl(link)
	if err != nil {
		return nil, err
	}

	account := &accountWithTimeStamp{
		createAt:  time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			link:     link,
		},
	}

	if account.password == "" {
		account.generatePassword(12)
	}

	return account, nil
}

func validateUrl(link string) error {
	_, err := url.ParseRequestURI(link)

	if err != nil {
		return errors.New("Неверный формат url")
	}
	return nil
}

func validateLogin(login string) error {
	if login == "" {
		return errors.New("Логин не может быть пустым")
	}

	return nil
}
