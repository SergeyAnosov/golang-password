package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(int(len(letterRunes)))]
	}
	acc.password = string(res)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("error. Login is empty")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	tempAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if tempAcc.password == "" {
		tempAcc.generatePassword(12)
	}

	return tempAcc, nil
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("error. Login is empty")
// 	}

// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("invalid URL")
// 	}

// 	tempAcc := &account{
// 		url:      urlString,
// 		login:    login,
// 		password: password,
// 	}

// 	if tempAcc.password == "" {
// 		tempAcc.generatePassword(12)
// 	}

// 	return tempAcc, nil
// }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUYWXZ1234567890-*!")

func main() {

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	if password == "" {
		fmt.Println("Пароль будет сгенерирован автоматически")
	}
	url := promptData("Введите URL: ")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Print("Не верный формат URL или LOGIN")
		return
	}

	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
