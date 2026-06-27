package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jihun")
	fmt.Println(account)
}
