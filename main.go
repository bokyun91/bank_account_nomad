package main

import (
	"fmt"

	"github.com/bokyun91/back_account_nomad/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
