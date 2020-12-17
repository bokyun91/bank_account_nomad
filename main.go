package main

import (
	"fmt"

	"github.com/bokyun91/back_account_nomad/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
