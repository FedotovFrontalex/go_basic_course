package account

import (
	"demo/app-4/print"
	"demo/app-4/validation"
	"math/rand/v2"
	"time"
)

var passwordChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Link      string    `json:"url"`
	CreateAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Print() {
	print.Data("Логин: ", acc.Login)
	print.Data("Пароль: ", acc.Password)
	print.Data("Ссылка: ", acc.Link)
	print.Data("Создан: ", acc.CreateAt)
	print.Data("Обновлен: ", acc.UpdatedAt)
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	lenghtChars := len(passwordChars)

	for index := range pass {
		pass[index] = passwordChars[rand.IntN(lenghtChars)]
	}

	acc.Password = string(pass)
}

func NewAccount(login string, password string, link string) (*Account, error) {
	var err error

	err = validation.ValidateLogin(login)
	if err != nil {
		return nil, err
	}

	err = validation.ValidateUrl(link)
	if err != nil {
		return nil, err
	}

	account := &Account{
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Link:      link,
	}

	if account.Password == "" {
		account.generatePassword(12)
	}

	return account, nil
}
