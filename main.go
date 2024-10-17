package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func menuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println("Счётчик вызова основного меню: ", i)
	}
}

func main() {
	fmt.Println("__Менеджер парольей__")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env файл")
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
	counter := menuCounter()

Menu:
	for {
		counter()
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccountByURL(vault *account.VaultWithDb) {
	url := promptData("Введите URL: ")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска: ")
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
	url := promptData("Введите URL: ")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promptData(prompt ...string) string {
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
