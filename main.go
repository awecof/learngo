package main

import (
	"fmt"
	"learngo/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("jihun")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
