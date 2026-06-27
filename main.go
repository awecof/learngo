package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jihun")
	account.Deposit(10)
	fmt.Println(account)
}
