package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUYWXZ1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"ureatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login, acc.Password, acc.Url)
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(int(len(letterRunes)))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("error. Login is empty")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	tempAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}

	if tempAcc.Password == "" {
		tempAcc.generatePassword(12)
	}

	return tempAcc, nil
}
