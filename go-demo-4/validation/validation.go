package validation

import (
	"errors"
	"net/url"
)

func ValidateUrl(link string) error {
	_, err := url.ParseRequestURI(link)

	if err != nil {
		return errors.New("Неверный формат url")
	}
	return nil
}

func ValidateLogin(login string) error {
	if login == "" {
		return errors.New("Логин не может быть пустым")
	}

	return nil
}
