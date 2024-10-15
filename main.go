package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

var menu = map[string]string{}

func main() {	
	fmt.Println("__Менеджер парольей__")
	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		input := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch input {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин: "})
	password := promptData([]string{"Введите пароль: "})
	url := promptData([]string{"Введите URL: "})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL: "})
	accounts := vault.FindAccount(url)
	if len(accounts) == 0 {
		color.Red("аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL: "})
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}

}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
