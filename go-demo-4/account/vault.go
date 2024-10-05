package account

import (
	"demo/app-4/files"
	"demo/app-4/print"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	bytes, err := files.ReadFile("accounts.json")
	newVault := &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}

	if err != nil {
		return newVault
	}
	var vault Vault
	err = json.Unmarshal(bytes, &vault)
	if err != nil {
		print.Error(errors.New("json поврежден"))
		return newVault
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	vault.save()
}

func (vault *Vault) save() {
	data, err := vault.toBytes()
	if err != nil {
		print.Error(err)
		return
	}
	files.WriteFile(data, "accounts.json")
}

func (vault *Vault) toBytes() ([]byte, error) {
	bytes, err := json.Marshal(vault)
	if err != nil {
		return nil, errors.New("Не удалосьпреобразовать в json")
	}
	return bytes, nil
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var res []Account
	for _, account := range vault.Accounts {
		if strings.Contains(account.Link, url) {
			res = append(res, account)
		}
	}

	return res
}

func (vault *Vault) DeleteAccountByUrl(url string) error {
		accounts := []Account{}

		for _, account := range vault.Accounts {
				if account.Link != url {
						accounts = append(accounts, account)
				}
		}

		if len(accounts) == len(vault.Accounts) {
				return errors.New("Аккаунт не найден")
		}

		vault.Accounts = accounts
		vault.UpdatedAt = time.Now()
		vault.save()
		return nil
}
