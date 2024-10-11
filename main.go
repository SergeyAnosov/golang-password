package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер парольей__")
	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		input := printMenu()
		switch input {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

func printMenu() int {
	var input int
	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&input)
	return input
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL: ")
	accounts := vault.FindAccount(url)
	if len(accounts) == 0 {
		color.Red("аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL: ")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найденоа")
	}

}
