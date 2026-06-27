package main

import (
	"fmt"
	"learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	err := dictionary.Add("first", "hello")
	if err == nil {
		fmt.Println("that word was added")
	} else {
		fmt.Println(err)
	}
	err2 := dictionary.Add("first", "hello")
	if err2 == nil {
		fmt.Println("that word was added")
	} else {
		fmt.Println(err2)
	}
	fmt.Println(dictionary)
}
