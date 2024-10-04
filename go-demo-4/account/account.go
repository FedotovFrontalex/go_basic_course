package account

import (
	"fmt"
	"math/rand/v2"
	"time"
	"demo/app-4/validation"
	"github.com/fatih/color"
)

var passwordChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

type Account struct {
	login    string
	password string
	link     string
}

type AccountWithTimeStamp struct { 
	Account
	createAt  time.Time
	updatedAt time.Time
}

func (acc *Account) Print() {
		kc := color.New(color.FgCyan)
		vc := color.New(color.FgBlue).Add(color.Bold)

		kc.Print("Логин - ")
		vc.Println(acc.login)

		kc.Print("Пароль - ")
		vc.Println(acc.password)

		kc.Print("Ссылка - ")
		vc.Println(acc.link)
}

func (acc *AccountWithTimeStamp) Print() {
	fmt.Printf("Создан: %s\n", acc.createAt)
	fmt.Printf("Обновлен: %s\n", acc.updatedAt)
	acc.Account.Print()
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	lenghtChars := len(passwordChars)

	for index := range pass {
		pass[index] = passwordChars[rand.IntN(lenghtChars)]
	}

	acc.password = string(pass)
}

func NewAccountWithTimeStamp(login string, password string, link string) (*AccountWithTimeStamp, error) {
	var err error

	err = validation.ValidateLogin(login)
	if err != nil {
		return nil, err
	}

	err = validation.ValidateUrl(link)
	if err != nil {
		return nil, err
	}

	account := &AccountWithTimeStamp{
		createAt:  time.Now(),
		updatedAt: time.Now(),
		Account: Account{
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
