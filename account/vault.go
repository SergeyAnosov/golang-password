package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) FindAccount(str string) []Account {
	var result []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, str)
		if isMatched {
			result = append(result, account)
		}
	}
	return result
}

func (vault *Vault) DeleteAccount(str string) bool {
	var result []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, str)
		if !isMatched {
			result = append(result, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = result
	vault.save()
	return isDeleted
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
