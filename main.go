package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("__Менеджер парольей__")
	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		})
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
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

func findAccountByURL(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL: "})
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)	
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин для поиска: "})
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, login)
	})
	outputResult(&accounts)	
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("аккаунтов не найдено")
	}
	for _, account := range *accounts {
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