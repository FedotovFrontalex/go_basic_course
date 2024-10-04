package main

import (
	"demo/app-4/account"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	link := promptData("Введите url: ")

	account1, err := account.NewAccountWithTimeStamp(login, password, link)

	if err != nil {	
			c := color.New(color.FgRed)
		c.Println(err)
		return
	}

	account1.Print()
}

func promptData(text string) string {
	var res string
	color.Cyan(text)
	fmt.Scanln(&res)
	return res
}
