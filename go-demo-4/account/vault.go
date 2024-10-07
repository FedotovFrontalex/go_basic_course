package account

import (
	"demo/app-4/cry"
	"demo/app-4/print"
	"encoding/json"
	"errors"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DB interface {
	Write([]byte)
	Read() ([]byte, error)
}

type VaultWithDb struct {
	Vault
	db  DB
	cry *cry.Crypto
}

func NewVault(db DB, cry *cry.Crypto) *VaultWithDb {
	encBytes, err := db.Read()

	newVault := &VaultWithDb{
		Vault: Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		},
		db:  db,
		cry: cry,
	}

	if err != nil {
		return newVault
	}
	var vault Vault
	bytes := cry.Decrypt(encBytes)
	err = json.Unmarshal(bytes, &vault)
	if err != nil {
		print.Error(errors.New("json поврежден"))
		return newVault
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		cry:   cry,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	vault.save()
}

func (vault *VaultWithDb) save() {
	data, err := vault.toBytes()
	if err != nil {
		print.Error(err)
		return
	}
	encData := vault.cry.Encrypt(data)
	vault.db.Write(encData)
}

func (vault *Vault) toBytes() ([]byte, error) {
	bytes, err := json.Marshal(vault)
	if err != nil {
		return nil, errors.New("Не удалосьпреобразовать в json")
	}
	return bytes, nil
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var res []Account
	for _, account := range vault.Accounts {
		if checker(account, str) {
			res = append(res, account)
		}
	}

	return res
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) error {
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
