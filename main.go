package main

import (
	"fmt"
	"demo/password/account"
	"demo/password/files"
)



func main() {

	files.ReadFile()
	files.WriteFile()
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	if password == "" {
		fmt.Println("Пароль будет сгенерирован автоматически")
	}
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Print("Не верный формат URL или LOGIN")
		return
	}

	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
